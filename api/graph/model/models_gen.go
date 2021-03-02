// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type CreateRoomInput struct {
	Name string `json:"name"`
}

type Message struct {
	ID        string    `json:"id"`
	User      *User     `json:"user"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

type RoomDetail struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Messages []*Message `json:"messages"`
}

type RoomSummary struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}
