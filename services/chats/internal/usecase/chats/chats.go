package chats

import (
	"context"

	"github.com/AlexeyTarasov77/messanger.chats/internal/dto"
	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/AlexeyTarasov77/messanger.chats/internal/gateways"
)

// UseCase -.
type UseCase struct {
	chatsRepo gateways.ChatsRepo
}

// New -.
func New(chatsRepo gateways.ChatsRepo) *UseCase {
	return &UseCase{
		chatsRepo: chatsRepo,
	}
}
func (uc *UseCase) CreatePersonalChat(ctx context.Context, dto dto.CreatePersonalChat) (*entity.PersonalChat, error)
func (uc *UseCase) CreateGroupChat(ctx context.Context, dto dto.CreateGroupChat) (*entity.PersonalChat, error)
