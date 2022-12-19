package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var Env Config

type Config struct {
	API_PORT string `json:"API_PORT"`

	MYSQL_DRIVER_NAME string `json:"MYSQL_DRIVER_NAME"`
	MYSQL_USERNAME    string `json:"MYSQL_USERNAME"`
	MYSQL_PASSWORD    string `json:"MYSQL_PASSWORD"`
	MYSQL_HOSTNAME    string `json:"MYSQL_HOSTNAME"`
	MYSQL_DB_NAME     string `json:"MYSQL_DB_NAME"`

	REDIS_ADDRESS  string `json:"REDIS_ADDRESS"`
	REDIS_PASSWORD string `json:"REDIS_PASSWORD"`
	REDIS_DB_NUM   int    `json:"REDIS_DB_NUM"`
}

func ConfigInit() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()
	viper.GetViper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper.ReadInConfig() => ConfigInit function failed. :: ", err)
		panic(err)
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		fmt.Println("viper.Unmarshal(&config) => ConfigInit function failed. :: ", err)
		panic(err)
	}

}
