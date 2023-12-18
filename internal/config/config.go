package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port               string
	APIKey             string
	PostgresHost       string
	PostgresDBName     string
	PostgresDBPort     string
	PostgresDBUser     string
	PostgresDBPassword string
	PostgresSSLMode    string
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
		Port:               v.GetString("PORT"),
		APIKey:             v.GetString("API_KEY"),
		PostgresHost:       v.GetString("POSTGRES_HOST"),
		PostgresDBName:     v.GetString("POSTGRES_DB"),
		PostgresDBPort:     v.GetString("POSTGRES_PORT"),
		PostgresDBUser:     v.GetString("POSTGRES_USER"),
		PostgresDBPassword: v.GetString("POSTGRES_PASSWORD"),
		PostgresSSLMode:    v.GetString("POSTGRES_SSL_MODE"),
	}
}
