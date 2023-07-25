package util

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func Init() {
	viper.AddConfigPath(".")
	
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config.toml")
	viper.SetConfigType("toml")
	
	Config = viper.GetViper()
	
	err := Config.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	InitLog()
}
