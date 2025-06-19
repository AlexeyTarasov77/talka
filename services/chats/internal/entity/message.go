package entity

import "time"

type Message struct {
	ID        int
	ChatId    int `json:"chat_id"`
	Text      string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
