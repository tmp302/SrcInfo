package Read

import (
	"github.com/spf13/viper"
)

func ReadConfig(targetInfo string) string{
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil{panic(err)}

	configInfo := viper.GetString(targetInfo)
	return configInfo
}