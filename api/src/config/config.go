package config

import (
	"log"
	"os"

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
	Port         string `default:"8080"`
	GoEnv        string `split_words:"true"`
	SessionKey   string `split_words:"true"`
	FrontBaseURL string `split_words:"true"`
	LogLevel     string `split_words:"true"`
	Db           DB
	Redis        Redis
	Twitter      Twitter
}

var Conf Config

func init() {
	if err := envconfig.Process("", &Conf); err != nil {
		log.Fatal("failed to read env variables,  err:", err)
	}
}

func IsDev() bool {
	return os.Getenv("GO_ENV") == "development"
}

func IsTest() bool {
	return os.Getenv("GO_ENV") == "test"
}

func IsProd() bool {
	return os.Getenv("GO_ENV") == "production"
}
