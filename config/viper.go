package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	settings := viper.New()

	settings.SetConfigName("config")
	settings.SetConfigType("json")
	settings.AddConfigPath("./")

	err := settings.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("viper not loading error %w \n", err))
	}

	return settings
}
