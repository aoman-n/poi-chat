package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Twitter struct {
	CallbackURI string `split_words:"true"`
	APIKey      string `split_words:"true"`
	SecretKey   string `split_words:"true"`
}

type DB struct {
	User     string `split_words:"true"`
	Name     string `split_words:"true"`
	Password string `split_words:"true"`
	Host     string `split_words:"true"`
	Port     string `split_words:"true"`
}

type Redis struct {
	URL      string `split_words:"true"`
	Passowrd string `split_words:"true"`
}

type Config struct {
	Port       string `default:"8080"`
	GoEnv      string `split_words:"true"`
	SessionKey string `split_words:"true"`
	Db         DB
	Redis      Redis
	Twitter    Twitter
}

var Conf Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file, err:", err)
	}

	err = envconfig.Process("", &Conf)
	if err != nil {
		log.Fatal("failed to read env variables,  err:", err)
	}
}
