package configuration

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/jaevor/go-nanoid"
	"github.com/spf13/viper"
)

const (
	CFG_APPLICATION_NAME       = "application.name"
	CFG_CONFIGURATION_PROFILES = "configuration.profiles"
)

var idGentor func() string

// InitializeAppConfig read and set default config values
func InitializeAppConfig() {
	slog.Info("InitializeAppConfig")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetDefault(CFG_APPLICATION_NAME, "GoApp")
	viper.SetDefault(CFG_CONFIGURATION_PROFILES, []string{})

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}

	profiles := viper.GetStringSlice(CFG_CONFIGURATION_PROFILES)
	for _, p := range profiles {
		profileCfg := "config." + p
		viper.SetConfigName(profileCfg)
		fmt.Printf("Loading %v\n", profileCfg)
		err := viper.MergeInConfig()
		if err != nil {
			slog.Info("MergeInConfig", "error", err)
		}
	}

	idGentor, _ = nanoid.Standard(21)
	configuredInstanceID := viper.GetString("application.instance")
	if configuredInstanceID == "" {
		id := idGentor()
		viper.Set("application.instance", id)
	}
}
