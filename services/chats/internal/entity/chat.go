package entity

import "time"

type ChatType string

var (
	PersonalChatType ChatType = "private-chat"
	GroupChatType    ChatType = "group-chat"

	ChatTypes = []ChatType{PersonalChatType, GroupChatType}
)

type (
	Chat interface {
		GetID() int
		GetTyp() ChatType
		GetCreatedAt() time.Time
		GetImageURL() string
		GetLastMsgText() string
		GetLastMsgDate() time.Time
	}
	BaseChat struct {
		ID          int
		Typ         ChatType
		CreatedAt   time.Time `json:"created_at"`
		ImageURL    string    `json:"image_url"`
		LastMsgText string    `json:"last_msg_text"`
		LastMsgDate time.Time `json:"last_msg_date"`
	}

	PersonalChat struct {
		Chat
		FromUserId int
		ToUserId   int
	}

	GroupChat struct {
		Chat
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

func (c *BaseChat) GetID() int {
	return c.ID
}

func (c *BaseChat) GetTyp() ChatType {
	return c.Typ
}

func (c *BaseChat) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *BaseChat) GetImageURL() string {
	return c.ImageURL
}

func (c *BaseChat) GetLastMsgText() string {
	return c.LastMsgText
}
func (c *BaseChat) GetLastMsgDate() time.Time {
	return c.LastMsgDate
}
