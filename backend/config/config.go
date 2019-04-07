package config

import (
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

func Init(path string) {
	var err error
	config = viper.New()
	config.SetConfigName("common")
	config.AddConfigPath(path)
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
