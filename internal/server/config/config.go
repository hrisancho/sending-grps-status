// SERVER Config
package config

import (
	"github.com/spf13/viper"
)

const (
	LoggerFormat = "${time} ${locals:requestid} ${status} - ${method} ${path}​\n"
)

type Config struct {
	// ServerAddr - адрес HTTP сервера (по умолчанию: 127.0.0.1:8080)
	ServerAddr string `mapstructure:"SERVER_ADDR"`

	// Logger settings
	Logger_fmt string `mapstructure:"LOGGER_FMT"`
}

func LoadConfig(path string) (config Config, err error) {
	// Config default values
	viper.SetDefault("SERVER_ADDR", "127.0.0.1:8080")
	viper.SetDefault("LOGGER_FMT", LoggerFormat)

	viper.AddConfigPath(path)
	viper.SetConfigName("main")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
