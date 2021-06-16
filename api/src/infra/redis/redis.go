package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/laster18/poi/api/src/config"
)

type Client = redis.Client

var Nil = redis.Nil

func New() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.URL,
		Password: config.Conf.Redis.Passowrd,
		DB:       config.Conf.Redis.Db,
	})

	context := context.Background()

	_, err := client.Ping(context).Result()

	if err != nil {
		log.Fatal("failed to connect redis", err)
	}

	log.Print("success to connect redis")

	return client
}
