package main

import (
	"fmt"
	"os"

	"github.com/laster18/poi/api/src/infra/server"
)

func main() {
	fmt.Printf(
		"TWITTER_API_KEY: %s\nTWITTER_SECRET_KEY: %s\nSESSION_KEY: %s\n",
		os.Getenv("TWITTER_API_KEY"),
		os.Getenv("TWITTER_SECRET_KEY"),
		os.Getenv("SESSION_KEY"),
	)

	server.Init()
}
