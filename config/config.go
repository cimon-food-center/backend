package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	AppPort string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
}

func InitConfig() (*config, error) {
	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

	config := new(config)
	config.AppPort = viper.GetString("app.port")
	config.DBHost = viper.GetString("db.host")
	config.DBPort = viper.GetString("db.port")
	config.DBUser = viper.GetString("db.user")
	config.DBPass = viper.GetString("db.pass")
	config.DBName = viper.GetString("db.name")

	return config, nil
}
