package main

import (
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/route53"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apiextensions"
	core "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/kustomize"
	meta "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	networking "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/networking/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"golang.org/x/crypto/bcrypt"
)

// Regular expression used to replace the kubeconfig token.
var re *regexp.Regexp = regexp.MustCompile(`(?m)^(\s*token:\s)\w*$`)

// Helper function to get an optional config or a default value.
func getOptionalConfig(cfg *config.Config, key string, def string) string {
	val := cfg.Get(key)
	if val == "" {
		val = def
	}
	return val
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Manage optional configuration.
		cfg := config.New(ctx, "")
		stack := ctx.Stack()
		cluster := getOptionalConfig(cfg, "clusterName", "enduro-sdps-"+stack)
		subdomain := getOptionalConfig(cfg, "route53Subdomain", stack+".sdps")
		zone := getOptionalConfig(cfg, "route53HostedZone", "artefactual.com")
		acmeServer := getOptionalConfig(cfg, "acmeServer", "https://acme-staging-v02.api.letsencrypt.org/directory")

		// Get DigitalOcean Kubernetes versions.
		k8sVersions, err := digitalocean.GetKubernetesVersions(ctx, nil)
		if err != nil {
			return err
		}

		// Create DigitalOcean Kubernetes cluster.
		k8sCluster, err := digitalocean.NewKubernetesCluster(ctx, "k8s-cluster",
			&digitalocean.KubernetesClusterArgs{
				Name:    pulumi.String(cluster),
				Region:  digitalocean.RegionNYC3,
				Version: pulumi.String(k8sVersions.LatestVersion),
				NodePool: &digitalocean.KubernetesClusterNodePoolArgs{
					Name:      pulumi.String(cluster + "-pool"),
					Size:      pulumi.String("s-4vcpu-8gb"),
					NodeCount: pulumi.Int(1),
				},
			},
		)
		if err != nil {
			return err
		}

		// Generate a non expiring kubeconfig for the cluster.
		kubeconfig := pulumi.All(
			k8sCluster.KubeConfigs.Index(pulumi.Int(0)).RawConfig().Elem(),
			cfg.RequireSecret("doToken"),
		).ApplyT(func(args []interface{}) string {
			return re.ReplaceAllString(args[0].(string), "${1}"+args[1].(string))
		}).(pulumi.StringOutput)

		// Create Kubernetes cluster provider with "sdps" as default namespace.
		k8sProvider, err := kubernetes.NewProvider(ctx, "k8s-provider",
			&kubernetes.ProviderArgs{
				Kubeconfig: kubeconfig,
				Namespace:  pulumi.StringPtr("sdps"),
			},
		)
		if err != nil {
			return err
		}

		// Create ingress-nginx namespace.
		nginxNS, err := core.NewNamespace(ctx, "nginx-ns",
			&core.NamespaceArgs{
				Metadata: &meta.ObjectMetaArgs{
					Name: pulumi.String("ingress-nginx"),
				},
			},
			pulumi.Provider(k8sProvider),
		)
		if err != nil {
			return err
		}

		// Install ingress-nginx Helm chart.
		nginxCtl, err := helm.NewRelease(ctx, "nginx-helm",
			&helm.ReleaseArgs{
				Chart:   pulumi.String("ingress-nginx"),
				Version: pulumi.String("4.1.4"),
				RepositoryOpts: &helm.RepositoryOptsArgs{
					Repo: pulumi.String("https://kubernetes.github.io/ingress-nginx"),
				},
				Namespace: nginxNS.Metadata.Name(),
				Values: pulumi.Map{
					"controller": pulumi.Map{
						"publishService": pulumi.Map{
							"enabled": pulumi.Bool(true),
						},
					},
				},
			},
			pulumi.Provider(k8sProvider),
		)
		if err != nil {
			return err
		}

		// Create cert-manager namespace.
		certNS, err := core.NewNamespace(ctx, "cert-manager-ns",
			&core.NamespaceArgs{
				Metadata: &meta.ObjectMetaArgs{
					Name: pulumi.String("cert-manager"),
				},
			},
			pulumi.Provider(k8sProvider),
		)
		if err != nil {
			return err
		}

		// Install cert-manager Helm chart.
		certMan, err := helm.NewRelease(ctx, "cert-manager-helm",
			&helm.ReleaseArgs{
				Chart:   pulumi.String("cert-manager"),
				Version: pulumi.String("1.8.1"),
				RepositoryOpts: &helm.RepositoryOptsArgs{
					Repo: pulumi.String("https://charts.jetstack.io"),
				},
				Namespace: certNS.Metadata.Name(),
				Values: pulumi.Map{
					"installCRDs": pulumi.Bool(true),
				},
			},
			pulumi.Provider(k8sProvider),
		)
		if err != nil {
			return err
		}

		// Configure default Docker images.
		crUrl := "registry.digitalocean.com"
		images := map[string]pulumi.Output{
			"enduro":            pulumi.ToOutput(crUrl + "/artefactual/enduro"),
			"enduro-a3m-worker": pulumi.ToOutput(crUrl + "/artefactual/enduro-a3m-worker"),
			"enduro-dashboard":  pulumi.ToOutput(crUrl + "/artefactual/enduro-dashboard"),
		}

		// Build, publish and update Docker images.
		if cfg.GetBool("buildImages") {
			err = buildAndPublishImages(ctx, crUrl, cfg.RequireSecret("doToken"), images)
			if err != nil {
				return err
			}
		}

		// Generate DigitalOcean container registry Docker config.
		crDockerConfig := cfg.RequireSecret("doToken").ApplyT(
			func(token string) string {
				return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(
					"{\"auths\":{\"%s\":{\"auth\":\"%s\"}}}",
					crUrl,
					base64.StdEncoding.EncodeToString([]byte(token+":"+token)),
				)))
			},
		).(pulumi.StringOutput)

		// Generate container registry credentials image pull secret.
		crSecret, err := core.NewSecret(ctx, "cr-secret",
			&core.SecretArgs{
				Metadata: &meta.ObjectMetaArgs{
					Name: pulumi.String("cr-secret"),
				},
				Data: pulumi.StringMap{
					".dockerconfigjson": crDockerConfig,
				},
				Type: pulumi.String("kubernetes.io/dockerconfigjson"),
			},
			pulumi.Provider(k8sProvider),
		)
		if err != nil {
			return err
		}

		// Apply Kubernetes base Kustomization, with the following transformations:
		// - Change Docker images to the ones from the DO CR.
		// - Add imagePullSecrets with the CR credentials secret.
		// - Set enduro-a3m replicas to 3.
		secrets := []map[string]interface{}{{"name": crSecret.Metadata.Name()}}
		k8sKustomize, err := kustomize.NewDirectory(ctx, "k8s-kustomize",
			kustomize.DirectoryArgs{
				Directory: pulumi.String("../kube/base"),
				Transformations: []yaml.Transformation{
					func(state map[string]interface{}, opts ...pulumi.ResourceOption) {
						name := state["metadata"].(map[string]interface{})["name"]
						if state["kind"] == "Deployment" && name == "enduro" {
							template := state["spec"].(map[string]interface{})["template"]
							templateSpec := template.(map[string]interface{})["spec"]
							containers := templateSpec.(map[string]interface{})["containers"]
							container := containers.([]interface{})[0]
							container.(map[string]interface{})["image"] = images["enduro"]
							templateSpec.(map[string]interface{})["imagePullSecrets"] = secrets
						} else if state["kind"] == "Deployment" && name == "enduro-dashboard" {
							template := state["spec"].(map[string]interface{})["template"]
							templateSpec := template.(map[string]interface{})["spec"]
							containers := templateSpec.(map[string]interface{})["containers"]
							container := containers.([]interface{})[0]
							container.(map[string]interface{})["image"] = images["enduro-dashboard"]
							templateSpec.(map[string]interface{})["imagePullSecrets"] = secrets
						} else if state["kind"] == "StatefulSet" && name == "enduro-a3m" {
							template := state["spec"].(map[string]interface{})["template"]
							templateSpec := template.(map[string]interface{})["spec"]
							containers := templateSpec.(map[string]interface{})["containers"]
							container := containers.([]interface{})[0]
							container.(map[string]interface{})["image"] = images["enduro-a3m-worker"]
							templateSpec.(map[string]interface{})["imagePullSecrets"] = secrets
							state["spec"].(map[string]interface{})["replicas"] = 3
						}
					},
				},
			},
			pulumi.Provider(k8sProvider),
		)
		if err != nil {
			return err
		}

		// Generate basic auth hash when username and password resolve.
		basicAuth := pulumi.All(
			cfg.RequireSecret("basicAuthUsername"),
			cfg.RequireSecret("basicAuthPassword"),
		).ApplyT(func(args []interface{}) (string, error) {
			username := args[0].(string)
			password := args[1].(string)
			bcryptPass, err := bcrypt.GenerateFromPassword(
				[]byte(password), bcrypt.DefaultCost,
			)
			if err != nil {
				return "", err
			}
			return base64.StdEncoding.EncodeToString([]byte(
				username + ":" + string(bcryptPass[:]),
			)), nil
		}).(pulumi.StringOutput)

		// Create basic auth secret.
		_, err = core.NewSecret(ctx, "basic-auth",
			&core.SecretArgs{
				Metadata: &meta.ObjectMetaArgs{
					Name: pulumi.String("basic-auth"),
				},
				Data: pulumi.StringMap{
					"auth": basicAuth,
				},
				Type: pulumi.String("Opaque"),
			},
			pulumi.Provider(k8sProvider),
		)
		if err != nil {
			return err
		}

		// Define endpoints for public services.
		type Endpoint struct {
			Name    string
			Service string
			Port    int
		}
		endpoints := []Endpoint{
			{Name: "enduro", Service: "enduro-dashboard", Port: 80},
			{Name: "minio", Service: "minio", Port: 9001},
			{Name: "temporal", Service: "temporal-ui", Port: 8080},
			{Name: "opensearch", Service: "opensearch-dashboards", Port: 5601},
		}

		// Generate ingress rules and TLS hosts for the endpoints.
		var hosts pulumi.StringArray
		var ingressRules networking.IngressRuleArray
		for _, endpoint := range endpoints {
			host := endpoint.Name + "." + subdomain + "." + zone
			ingressRule := &networking.IngressRuleArgs{
				Host: pulumi.String(host),
				Http: &networking.HTTPIngressRuleValueArgs{
					Paths: networking.HTTPIngressPathArray{
						&networking.HTTPIngressPathArgs{
							Path:     pulumi.String("/"),
							PathType: pulumi.String("Prefix"),
							Backend: &networking.IngressBackendArgs{
								Service: &networking.IngressServiceBackendArgs{
									Name: pulumi.String(endpoint.Service),
									Port: &networking.ServiceBackendPortArgs{
										Number: pulumi.Int(endpoint.Port),
									},
								},
							},
						},
					},
				},
			}
			ingressRules = append(ingressRules, ingressRule)
			hosts = append(hosts, pulumi.String(host))
		}

		// Create ingress.
		ingress, err := networking.NewIngress(ctx, "ingress",
			&networking.IngressArgs{
				Metadata: &meta.ObjectMetaArgs{
					Name: pulumi.String("ingress"),
					Annotations: pulumi.StringMap{
						"nginx.ingress.kubernetes.io/auth-type":       pulumi.String("basic"),
						"nginx.ingress.kubernetes.io/auth-secret":     pulumi.String("basic-auth"),
						"nginx.ingress.kubernetes.io/auth-realm":      pulumi.String("Authentication required!"),
						"nginx.ingress.kubernetes.io/proxy-body-size": pulumi.String("0"),
						"cert-manager.io/cluster-issuer":              pulumi.String("cert-issuer"),
					},
				},
				Spec: &networking.IngressSpecArgs{
					IngressClassName: pulumi.String("nginx"),
					Tls: networking.IngressTLSArray{
						&networking.IngressTLSArgs{
							Hosts:      hosts,
							SecretName: pulumi.String("acme-cert"),
						},
					},
					Rules: ingressRules,
				},
			},
			pulumi.DependsOn([]pulumi.Resource{nginxCtl, k8sKustomize}),
			pulumi.Provider(k8sProvider),
		)
		if err != nil {
			return err
		}

		// Get AWS Route 53 zone.
		route53Zone, err := route53.LookupZone(
			ctx, &route53.LookupZoneArgs{Name: pulumi.StringRef(zone)},
		)
		if err != nil {
			return err
		}

		// Create AWS Route 53 records for the endpoints.
		var dnsResources []pulumi.Resource
		ingressIp := ingress.Status.LoadBalancer().Ingress().Index(pulumi.Int(0)).Ip().Elem()
		for _, endpoint := range endpoints {
			dnsResource, err := route53.NewRecord(ctx, endpoint.Name+"-dns",
				&route53.RecordArgs{
					ZoneId: pulumi.String(route53Zone.ZoneId),
					Name:   pulumi.String(endpoint.Name + "." + subdomain + "." + zone),
					Type:   pulumi.String("A"),
					Ttl:    pulumi.Int(300),
					Records: pulumi.StringArray{
						ingressIp,
					},
				},
			)
			if err != nil {
				return err
			}
			dnsResources = append(dnsResources, dnsResource)
		}

		// Create cert-manager cluster issuer.
		_, err = apiextensions.NewCustomResource(ctx, "cert-issuer",
			&apiextensions.CustomResourceArgs{
				ApiVersion: pulumi.String("cert-manager.io/v1"),
				Kind:       pulumi.String("ClusterIssuer"),
				Metadata: &meta.ObjectMetaArgs{
					Name:      pulumi.String("cert-issuer"),
					Namespace: pulumi.String("cert-manager"),
				},
				OtherFields: kubernetes.UntypedArgs{
					"spec": kubernetes.UntypedArgs{
						"acme": kubernetes.UntypedArgs{
							"email":  cfg.RequireSecret("acmeEmail"),
							"server": acmeServer,
							"privateKeySecretRef": kubernetes.UntypedArgs{
								"name": pulumi.String("acme-secret"),
							},
							"solvers": []kubernetes.UntypedArgs{{
								"http01": kubernetes.UntypedArgs{
									"ingress": kubernetes.UntypedArgs{
										"class": pulumi.String("nginx"),
									},
								},
							}},
						},
					},
				},
			},
			pulumi.DependsOn(append(dnsResources, certMan)),
			pulumi.Provider(k8sProvider),
		)
		if err != nil {
			return err
		}

		return nil
	})
}
