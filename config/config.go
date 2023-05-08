package config

//nolint:unused

import (
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type (
	Config struct {
		Logger   LoggerConf   `mapstructure:"logger"`
		Server   ServerConf   `mapstructure:"httpserver"`
		Database DatabaseConf `mapstructure:"database"`
		Rpc      RPC          `mapstructure:"rpcserver"`
	}

	LoggerConf struct {
		Level      zapcore.Level `mapstructure:"level" default:"debug"`
		Encoding   string        `mapstructure:"encoding" default:"console"`
		OutputPath []string      `mapstructure:"output" default:"stdout"`
	}

	ServerConf struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}

	DatabaseConf struct {
		Url            string `mapstructure:"dsn"`
		Type           string `mapstructure:"type"`
		Implementation string `mapstructure:"implementation" default:"sql"`
	}
	RPC struct {
		Host string `mapstructure:"host" default:"localhost"`
		Port int    `mapstructure:"port" default:"54323"`
	}
)

func NewConfig(path string) (Config, error) {
	viper.SetConfigFile(path)
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
