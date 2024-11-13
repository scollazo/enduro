package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/artefactual-sdps/temporal-activities/archiveextract"
	"github.com/artefactual-sdps/temporal-activities/bagcreate"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"go.artefactual.dev/tools/bucket"

	"github.com/artefactual-sdps/enduro/internal/a3m"
	"github.com/artefactual-sdps/enduro/internal/am"
	"github.com/artefactual-sdps/enduro/internal/api"
	"github.com/artefactual-sdps/enduro/internal/db"
	"github.com/artefactual-sdps/enduro/internal/event"
	"github.com/artefactual-sdps/enduro/internal/package_"
	"github.com/artefactual-sdps/enduro/internal/poststorage"
	"github.com/artefactual-sdps/enduro/internal/premis"
	"github.com/artefactual-sdps/enduro/internal/preprocessing"
	"github.com/artefactual-sdps/enduro/internal/pres"
	"github.com/artefactual-sdps/enduro/internal/storage"
	"github.com/artefactual-sdps/enduro/internal/telemetry"
	"github.com/artefactual-sdps/enduro/internal/temporal"
	"github.com/artefactual-sdps/enduro/internal/watcher"
)

type ConfigurationValidator interface {
	Validate() error
}

type Configuration struct {
	Debug       bool
	DebugListen string
	Verbosity   int

	A3m             a3m.Config
	AM              am.Config
	InternalAPI     api.Config
	API             api.Config
	BagIt           bagcreate.Config
	Database        db.Config
	Event           event.Config
	ExtractActivity archiveextract.Config
	Poststorage     []poststorage.Config
	Preprocessing   preprocessing.Config
	Preservation    pres.Config
	Storage         storage.Config
	Temporal        temporal.Config
	Upload          package_.UploadConfig
	Watcher         watcher.Config
	Telemetry       telemetry.Config
	ValidatePREMIS  premis.Config

	FailedSIPs bucket.Config
	FailedPIPs bucket.Config
}

func (c *Configuration) Validate() error {
	// TODO: should this validate all the fields in Configuration?
	return errors.Join(
		c.A3m.Validate(),
		c.API.Auth.Validate(),
		c.BagIt.Validate(),
		c.Preprocessing.Validate(),
		c.Upload.Validate(),
		c.ValidatePREMIS.Validate(),
	)
}

func Read(config *Configuration, configFile string) (found bool, configFileUsed string, err error) {
	v := viper.New()

	v.AddConfigPath(".")
	v.AddConfigPath("$HOME/.config/")
	v.AddConfigPath("/etc")
	v.SetConfigName("enduro")
	v.SetDefault("a3m.capacity", 1)
	v.SetDefault("a3m.processing", a3m.ProcessingDefault)
	v.SetDefault("am.capacity", 1)
	v.SetDefault("am.pollInterval", 10*time.Second)
	v.SetDefault("api.listen", "127.0.0.1:9000")
	v.SetDefault("debugListen", "127.0.0.1:9001")
	v.SetDefault("preservation.taskqueue", temporal.A3mWorkerTaskQueue)
	v.SetDefault("storage.taskqueue", temporal.GlobalTaskQueue)
	v.SetDefault("temporal.taskqueue", temporal.GlobalTaskQueue)
	v.SetDefault("upload.maxSize", 102400000)
	v.SetEnvPrefix("enduro")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if configFile != "" {
		v.SetConfigFile(configFile)
	}

	err = v.ReadInConfig()
	_, ok := err.(viper.ConfigFileNotFoundError)
	if !ok {
		found = true
	}
	if found && err != nil {
		return found, configFileUsed, fmt.Errorf("failed to read configuration file: %w", err)
	}

	decodeHookFunc := mapstructure.ComposeDecodeHookFunc(
		// These are the viper DecodeHookFunc defaults.
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
		// StringToUUIDHookFunc is a custom string to UUID decoder.
		stringToUUIDHookFunc(),
		stringToMapHookFunc(),
	)

	err = v.Unmarshal(config, viper.DecodeHook(decodeHookFunc))
	if err != nil {
		return found, configFileUsed, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	if err := config.Validate(); err != nil {
		return found, configFileUsed, fmt.Errorf("failed to validate the provided config: %w", err)
	}

	configFileUsed = v.ConfigFileUsed()

	if err := setCORSOriginEnv(config); err != nil {
		return found, configFileUsed, fmt.Errorf(
			"failed to set CORS Origin environment variable: %w", err,
		)
	}

	return found, configFileUsed, nil
}

// setCORSOriginEnv sets the CORS Origin environment variable needed by
// Goa-generated code for the API.
func setCORSOriginEnv(config *Configuration) error {
	if config.API.CORSOrigin == "" {
		// Default to the API URI to disallow all cross-origin requests.
		config.API.CORSOrigin = config.API.Listen
	}

	if err := os.Setenv("ENDURO_API_CORS_ORIGIN", config.API.CORSOrigin); err != nil {
		return err
	}

	return nil
}

// stringToUUIDHookFunc decodes a string to a uuid.UUID. Copied from
// https://github.com/go-saas/kit/blob/main/pkg/mapstructure/mapstructure.go
func stringToUUIDHookFunc() mapstructure.DecodeHookFunc {
	return func(f, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String || t != reflect.TypeOf(uuid.UUID{}) {
			return data, nil
		}

		return uuid.Parse(data.(string))
	}
}

// stringToMapHookFunc decodes a JSON string to a map[string][]string.
func stringToMapHookFunc() mapstructure.DecodeHookFunc {
	return func(f, t reflect.Type, data interface{}) (interface{}, error) {
		value := map[string][]string{}
		if f.Kind() != reflect.String || t != reflect.TypeOf(value) {
			return data, nil
		}

		if data.(string) != "" {
			if err := json.Unmarshal([]byte(data.(string)), &value); err != nil {
				return nil, err
			}
		}

		return value, nil
	}
}
