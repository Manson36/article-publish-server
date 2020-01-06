package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("/go/bin/config") //prod environment
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("config file read error, msg:%s", err.Error()))
	}

	Web.readConf()
	Postgres.readConf()
	Server.readConf()
	Redis.readConfig()
}
