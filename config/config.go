package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	API_PORT string `mapstructure:"API_PORT"`

	MYSQL_DRIVER_NAME string `mapstructure:"MYSQL_DRIVER_NAME"`
	MYSQL_USERNAME    string `mapstructure:"MYSQL_USERNAME"`
	MYSQL_PASSWORD    string `mapstructure:"MYSQL_PASSWORD"`
	MYSQL_HOSTNAME    string `mapstructure:"MYSQL_HOSTNAME"`
	MYSQL_DB_NAME     string `mapstructure:"MYSQL_DB_NAME"`

	REDIS_ADDRESS  string `mapstructure:"REDIS_ADDRESS"`
	REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`
	REDIS_DB_NUM   int    `mapstructure:"REDIS_DB_NUM"`
}

func ConfigInit() (config Config) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.GetViper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper.ReadInConfig() => ConfigInit function failed. :: ", err)
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("viper.Unmarshal(&config) => ConfigInit function failed. :: ", err)
		panic(err)
	}

	return
}
