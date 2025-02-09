# Enduro - Pulumi

Pulumi project to manage Enduro's infrastructure in DigitalOcean Kubernetes
clusters and AWS Route 53 records, with Let's Encrypt certificates and basic
authentication.

## Requirements

- [Go] (v1.18.1+)
- [Pulumi] (v3.35.3+)

To build new images, [Docker] (v18.09+) is also required and [BuildKit builds]
must be enabled.

## Set up

To install the Go dependencies, located in this directory, execute:

```
go mod download
```

## Login

The project backend is managed with the Pulumi service. To login, run:

```
pulumi login
```

Hit `Enter` to open the browser, select sign in with email and use the
credentials from LastPass.

## Project configuration

Use the [Pulumi CLI] `config set` and `config set-all` commands to configure.

### Required

- `doToken` **secret**: DigitalOcean API token.
- `digitalocean:token` **secret**: DigitalOcean API token.
- `acmeEmail` **secret**: ACME registration email.
- `aws:region`: AWS region.
- `aws:accessKey` **secret**: AWS access key.
- `aws:secretKey` **secret**: AWS secret key.
- `basicAuthUsername` **secret**: Basic authentication username.
- `basicAuthPassword` **secret**: Basic authentication password.

### Optional

- `acmeServer`: ACME server URL. Default: `https://acme-v02.api.letsencrypt.org/directory`
- `route53HostedZone`: AWS Route 53 hosted zone. Default: `artefactual.com`
- `route53Subdomain`: Route 53 records subdomain. Default: `<stack-name>.sdps`
- `clusterName`: DigitalOcean k8s cluster name. Default: `enduro-sdps-<stack-name>`
- `buildImages`: Whether to build new images or not. Default: `false`

### Example

```
pulumi config set-all \
  --secret doToken=abc123 \
  --secret digitalocean:token=abc123 \
  --plaintext acmeServer=https://acme-staging-v02.api.letsencrypt.org/directory \
  --secret acmeEmail=user@example.com \
  --plaintext aws:region=us-east-1 \
  --secret aws:accessKey=abc123 \
  --secret aws:secretKey=abc123 \
  --plaintext route53HostedZone=example.com \
  --plaintext route53Subdomain=sub.domain \
  --plaintext clusterName=my-sdps-cluster \
  --secret basicAuthUsername=abc123 \
  --secret basicAuthPassword=abc123 \
  --plaintext buildImages=true
```

## Create stack

```
pulumi stack init <stack-name>

pulumi config set-all ...

pulumi up
```

## Update stack

```
pulumi stack select <stack-name>

pulumi up --refresh
```

## Delete stack

```
pulumi stack select <stack-name>

pulumi destroy

pulumi stack rm <stack-name>
```

[go]: https://go.dev/doc/install
[pulumi]: https://www.pulumi.com/docs/get-started/install/
[docker]: https://docs.docker.com/get-docker/
[buildkit builds]: https://docs.docker.com/develop/develop-images/build_enhancements/
[pulumi cli]: https://www.pulumi.com/docs/reference/cli/
