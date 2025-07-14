package entity

import "time"

type (
	Membership struct {
		UserId           int       `json:"user_id"`
		GroupChatId      int       `json:"group_chat_id"`
		CreatedAt        time.Time `json:"created_at"`
		InvitationLinkId int       `json:"invitation_link_id"`
	}
	ChatJoinReq struct {
		GroupChatId      int
		FromUserId       int
		CreatedAt        time.Time `json:"created_at"`
		InvitationLinkId int       `json:"invitation_link_id"`
	}
)
