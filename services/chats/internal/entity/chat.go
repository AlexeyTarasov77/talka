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
		GetLastMsg() *Message
	}
	ChatWithMessages interface {
		Chat
		GetMessages() []Message
		SetMessages([]Message)
	}
	BaseChat struct {
		ID        int
		Typ       ChatType
		CreatedAt time.Time `json:"created_at"`
		ImageURL  string    `json:"image_url"`
		// LastMsg should be kept always up to date
		LastMsg *Message
		// Messages are retrieved only in special cases
		Messages []Message
	}

	PersonalChat struct {
		Chat
		FromUserId int
		ToUserId   int
	}

	GroupChat struct {
		Chat
		// OwnerId and Name create unique pair (same user can't have multiple chats with same name)
		OwnerId     int
		Name        string
		Description string // OPTIONAL
		// PrimaryLinkUrl is unique field which is used in chat url.
		// User can join to the group by using this link.
		// Link behaviour depends on the group type (public/private):
		// - PUBLIC:
		//   - must be provided by user and can be freely changed in the same way
		//   - user can immediately join group by using this link
		// - PRIVATE:
		//   - a secure value must be automatically generated, can be "revoked" which means regenerating value
		//   - user can only send a join request which requires approval to join by using this link
		PrimaryLinkUrl string
		// IsPublic indicates whether chat is public and anybody can join it or user can join only by sending request for approval
		// DEFAULT: false
		IsPublic bool
		Members  []User
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

func (c *BaseChat) GetLastMsg() *Message {
	return c.LastMsg
}

func (c *BaseChat) GetMessages() []Message {
	return c.Messages
}

func (c *BaseChat) SetMessages(messages []Message) {
	c.Messages = messages
}
