package entity

import "time"

type User struct {
	ID          int
	Username    string
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	AvatarUrl   string    `json:"avatar_url"`
	LastSeenAt  time.Time `json:"last_seen_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsActive    bool      `json:"is_active"`
	BirthDate   time.Time `json:"birth_date"`
}
