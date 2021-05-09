package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/db"
	"gorm.io/gorm"
)

func seedRoomsAndMessages(db *gorm.DB) {
	tx := db.Begin()
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			tx.Commit()
		}
	}()

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

	if err := tx.Create(&rooms).Error; err != nil {
		panic(err)
	}

	for _, room := range rooms {
		willMadeCount := 20
		messages := make([]domain.Message, willMadeCount)
		for i := range make([]int, willMadeCount) {
			msg := &domain.Message{
				UserUID:       fmt.Sprintf("user_%d", i),
				Body:          fmt.Sprintf("Room - %d message!", room.ID),
				UserName:      faker.Name(),
				UserAvatarURL: "https://pbs.twimg.com/profile_images/1130684542732230656/pW77OgPS_400x400.png",
				RoomID:        room.ID,
				CreatedAt:     time.Unix(faker.RandomUnixTime(), 0),
			}
			messages[i] = *msg
		}

		if err := tx.Create(&messages).Error; err != nil {
			log.Print("err:", err)
			panic(err)
		}

		fmt.Printf("createdRoom: %v\n", room)
		fmt.Printf("createdMessages: %v\n", messages)
	}
}

func seedMessages() {

}

func main() {
	db := db.NewDb()
	seedRoomsAndMessages(db)

}
