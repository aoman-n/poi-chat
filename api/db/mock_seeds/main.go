package main

import (
	"fmt"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infrastructure"
)

func seedRooms(db *infrastructure.Db) {
	rooms := []domain.Room{
		{
			Name:            "room1",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
		},
		{
			Name:            "room2",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
		},
		{
			Name:            "room3",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
		},
		{
			Name:            "room4",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
		},
		{
			Name:            "room5",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
		},
		{
			Name:            "room6",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
		},
		{
			Name:            "room7",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
		},
		{
			Name:            "room8",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
		},
		{
			Name:            "room9",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
		},
		{
			Name:            "room10",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
		},
	}

	db.Create(&rooms)

	for _, room := range rooms {
		fmt.Printf("created: %v", room)
	}
}

func main() {
	db := infrastructure.NewDb()
	seedRooms(db)
}
