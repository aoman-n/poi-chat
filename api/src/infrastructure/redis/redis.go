package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/laster18/poi/api/src/config"
)

func New(conf config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.URL,
		Password: conf.Passowrd,
		DB:       0,
	})

	context := context.Background()

	_, err := client.Ping(context).Result()

	if err != nil {
		log.Fatal("failed to connect redis", err)
	}

	log.Print("success to connect redis")

	return client
}
