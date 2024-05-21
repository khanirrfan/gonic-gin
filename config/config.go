package config

import (
	"log"

	"github.com/spf13/viper"
)

/*
*
Viper simplifies configuration management in Go applications
by providing a unified and flexible way to read from multiple
configuration sources. This helps in maintaining clean and
maintainable code, especially in complex applications with
multiple configuration needs
*/
var config *viper.Viper

var AppConfig *viper.Viper

func Init() {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err.Error())
	}
}

func GetConfig() *viper.Viper {
	return config
}
