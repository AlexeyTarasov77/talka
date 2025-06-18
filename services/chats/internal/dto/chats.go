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
)

func (payload *CreateGroupChat) MapToEntity() *entity.GroupChat {
	var chatMembers []entity.User
	if payload.MembersIds != nil && len(payload.MembersIds) > 0 {
		chatMembers = make([]entity.User, 0, len(payload.MembersIds))
		for _, id := range payload.MembersIds {
			chatMembers = append(chatMembers, entity.User{ID: id})
		}
	}
	return &entity.GroupChat{
		Chat:        &entity.BaseChat{Typ: entity.GroupChatType, ImageURL: payload.ImageURL},
		OwnerId:     payload.CurrentUserId,
		Name:        payload.Name,
		Description: payload.Description,
		Slug:        payload.Slug,
		Members:     chatMembers,
		IsPublic:    payload.IsPublic,
	}
}
