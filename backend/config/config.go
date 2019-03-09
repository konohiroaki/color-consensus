package config

import (
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

func Init() {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("common")
	config.AddConfigPath("backend/config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
