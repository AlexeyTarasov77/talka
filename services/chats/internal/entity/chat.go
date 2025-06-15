package entity

import "time"

type ChatType string

var (
	PrivateChatType ChatType = "private-chat"
	GroupChatType   ChatType = "group-chat"
)

type (
	Chat interface {
		GetID() int
		GetTyp() ChatType
		GetCreatedAt() time.Time
		GetImageURL() string
	}
	chatImpl struct {
		ID        int
		Typ       ChatType
		CreatedAt time.Time `json:"created_at"`
		ImageURL  string    `json:"image_url"`
	}

	PersonalChat struct {
		chatImpl
		FromUserId int
		ToUserId   int
	}

	GroupChat struct {
		chatImpl
		OwnerId     int
		Name        string
		Description string // OPTIONAL
		// Slug is unique field which is used in chat url.
		// When group is not public - url with this slug serves purpose of invitation link
		// Value should be automatically generated if not provided explicitly.
		// Slug can be revoked - in this case new slug should be automatically generated
		Slug string
		// IsPublic indicates whether chat is public and anybody can join it or user can join only by sending request for approval
		// DEFAULT: false
		IsPublic bool
	}
)
