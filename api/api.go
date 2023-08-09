package api

import "github.com/spf13/viper"

var SigningKey string

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	SigningKey = viper.GetString("SIGNING_KEY")
}
