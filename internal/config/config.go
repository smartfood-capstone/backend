package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	PORT                 int
	API_KEY              string
	POSTGRES_HOST        string
	POSTGRES_DB_NAME     string
	POSTGRES_DB_PORT     string
	POSTGRES_DB_USER     string
	POSTGRES_DB_PASSWORD string
}

func New() Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigName(".env")
	v.SetConfigType("env")

	v.AddConfigPath(".")
	v.AddConfigPath("../")
	v.AddConfigPath("../../")
	v.AddConfigPath("../../../../")

	err := v.ReadInConfig()
	if err != nil {
		logrus.Fatalf("error when reading config file: %s", err)
	}

	return Config{
		PORT:                 v.GetInt("PORT"),
		API_KEY:              v.GetString("API_KEY"),
		POSTGRES_HOST:        v.GetString("PORT"),
		POSTGRES_DB_NAME:     v.GetString("PORT"),
		POSTGRES_DB_PORT:     v.GetString("PORT"),
		POSTGRES_DB_USER:     v.GetString("PORT"),
		POSTGRES_DB_PASSWORD: v.GetString("PORT"),
	}
}
