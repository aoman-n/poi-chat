package main

import (
	"fmt"
	"time"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infrastructure/db"
	"gorm.io/gorm"
)

func seedRooms(db *gorm.DB) {
	rooms := []domain.Room{
		{
			ID:              1,
			Name:            "room1",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
			CreatedAt:       time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:              2,
			Name:            "room2",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
			CreatedAt:       time.Date(2021, 2, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:              3,
			Name:            "room3",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
			CreatedAt:       time.Date(2021, 2, 3, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:              4,
			Name:            "room4",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
			CreatedAt:       time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:              5,
			Name:            "room5",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
			CreatedAt:       time.Date(2021, 3, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:              6,
			Name:            "room6",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
			CreatedAt:       time.Date(2021, 1, 6, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:              7,
			Name:            "room7",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
			CreatedAt:       time.Date(2021, 3, 7, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:              8,
			Name:            "room8",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
			CreatedAt:       time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:              9,
			Name:            "room9",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
			CreatedAt:       time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:              10,
			Name:            "room10",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
			CreatedAt:       time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:              11,
			Name:            "room11",
			BackgroundURL:   "https://example.com/image.png",
			BackgroundColor: "#66cdaa",
			UserCount:       0,
			CreatedAt:       time.Date(2021, 2, 3, 0, 0, 0, 0, time.UTC),
		},
	}

	db.Create(&rooms)

	for _, room := range rooms {
		fmt.Printf("created: %v\n", room)
	}
}

func main() {
	db := db.NewDb()
	seedRooms(db)
}
