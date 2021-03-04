package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Environments struct {
	Port             string `default:"8080"`
	GoEnv            string `split_words:"true"`
	DbUser           string `split_words:"true"`
	DbName           string `split_words:"true"`
	DbPassword       string `split_words:"true"`
	DbHost           string `split_words:"true"`
	DbPort           string `split_words:"true"`
	RedisURL         string `split_words:"true"`
	RedisPassowrd    string `split_words:"true"`
	TwitterAPIKey    string `split_words:"true"`
	TwitterSecretKey string `split_words:"true"`
}

var Env Environments

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file, err:", err)
	}

	err = envconfig.Process("", &Env)
	if err != nil {
		log.Fatal("failed to read env variables,  err:", err)
	}
}
