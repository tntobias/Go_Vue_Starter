package config

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Port      string        `yaml:"port,omitempty"`
	Secret    string        `yaml:"secret,omitempty"`
	LogConfig LoggingConfig `yaml:"log_config,omitempty"`
}

// LoggingConfig specifies all the parameters needed for logging
type LoggingConfig struct {
	Level string `yaml:"level,omitempty"`
	File  string `yaml:"file,omitempty"`
}

func LoadConfig(cmd *cobra.Command) (*Config, error) {
	viper.SetConfigType("yaml")

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return nil, err
	}

	viper.SetEnvPrefix("PRJ")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("./")
		viper.AddConfigPath("/etc/myapp")
	}

	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
