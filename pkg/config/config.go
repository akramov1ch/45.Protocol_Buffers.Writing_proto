package config

import (
	"log"

	"github.com/spf13/viper"
)

type Conn struct {
	HTTPServerPort    string
	JSONRPCServerPort string
	DatabaseURL       string
}

var Conf Conn

func LoadConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	Conf.HTTPServerPort = viper.GetString("HTTP_SERVER_PORT")
	Conf.JSONRPCServerPort = viper.GetString("JSON_RPC_SERVER_PORT")
	Conf.DatabaseURL = viper.GetString("DATABASE_URL")
}
