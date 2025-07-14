package dto

import "github.com/AlexeyTarasov77/messanger.chats/internal/entity"

type (
	CreatePersonalChat struct {
		CurrentUserId int `json:"current_user_id"`
		WithUserId    int `json:"with_user_id"`
	}
	CreateGroupChat struct {
		CurrentUserId int    `json:"current_user_id"`
		ImageURL      string `json:"image_url"`
		MembersIds    []int  `json:"members_ids"`
		Name          string
		Description   string
		Slug          string
		IsPublic      bool `json:"is_public"`
	}
	JoinChat struct {
		CurrentUserId int `json:"current_user_id"`
		ChatId        int `json:"chat_id"`
	}
)

func (payload *CreateGroupChat) MapToEntity() *entity.GroupChat {
	return &entity.GroupChat{
		Chat:           &entity.BaseChat{Typ: entity.GroupChatType, ImageURL: payload.ImageURL},
		OwnerId:        payload.CurrentUserId,
		Name:           payload.Name,
		Description:    payload.Description,
		PrimaryLinkUrl: payload.Slug,
		IsPublic:       payload.IsPublic,
	}
}
