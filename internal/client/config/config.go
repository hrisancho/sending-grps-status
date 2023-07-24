// CLIENT Config
package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"reflect"
)

type Config struct {
	// ServerAddr - адрес сервера
	ServerAddr string `mapstructure:"SERVER_ADDR"`
	UUID       string `mapstructure:"UUID"`
}

func initDefaultConfig() (v *viper.Viper) {
	v = viper.New()
	v.SetDefault("SERVER_ADDR", "127.0.0.1:8080")
	v.SetDefault("UUID", "1412fe28-06b0-408b-832d-f2d5aa4792a5")
	return v
}

func loadConfigFile(v *viper.Viper, path string) (config Config, err error) {
	v.AddConfigPath(path)
	v.SetConfigName("main")
	v.SetConfigType("env")

	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	configReflectType := reflect.ValueOf(&config).Elem()
	configFieldsCount := configReflectType.NumField()

	err = v.Unmarshal(&config)
	if err != nil {
		return
	}

	for fieldInd := 0; fieldInd < configFieldsCount; fieldInd++ {
		configField := configReflectType.Field(fieldInd)

		if configField.Kind() != reflect.Struct {
			continue
		}

		err = v.Unmarshal(configField.Addr().Interface())
		if err != nil {
			return
		}
	}

	return
}

func loadConfigEnv(v *viper.Viper) (config Config, err error) {
	envNameList := envNameListByConfig(reflect.TypeOf(config))
	for _, envName := range envNameList {
		err = v.BindEnv(envName, envName)
		if err != nil {
			return
		}
	}

	err = v.Unmarshal(&config)
	return
}

func envNameListByConfig(configType reflect.Type) (envNameList []string) {
	configFieldsCount := configType.NumField()
	envNameList = make([]string, 0, configFieldsCount)

	for fieldInd := 0; fieldInd < configFieldsCount; fieldInd++ {
		configField := configType.Field(fieldInd)

		if configField.Type.Kind() == reflect.Struct {
			envNameList = append(envNameList, envNameListByConfig(configField.Type)...)
		}

		envNameList = append(envNameList, configField.Tag.Get("mapstructure"))
	}
	return
}
func LoadConfig(path string) (config Config, err error) {
	// Config default values
	conf := initDefaultConfig()

	if _, err = os.Stat("./config/main.env"); err == nil {
		log.Println("Loading config from file...")
		config, err = loadConfigFile(conf, "./config")
	} else {
		log.Println("Loading config from env...")
		config, err = loadConfigEnv(conf)
	}

	return config, err
}
