package entity

import "time"

type (
	Membership struct {
		UserId         int       `json:"user_id"`
		GroupChatId    int       `json:"group_chat_id"`
		CreatedAt      time.Time `json:"created_at"`
		JoinedByLinkId int       `json:"joined_by_link_id"`
	}
)
