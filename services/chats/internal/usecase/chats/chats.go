package chats

import (
	"context"
	"errors"

	"github.com/AlexeyTarasov77/messanger.chats/internal/dto"
	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/AlexeyTarasov77/messanger.chats/internal/gateways"
	"github.com/AlexeyTarasov77/messanger.chats/internal/gateways/storage"
)

// UseCase -.
type UseCase struct {
	chatsRepo gateways.ChatsRepo
	txManager gateways.TransactionsManager
	usersApi  gateways.UsersAPI
}

// New -.
func New(chatsRepo gateways.ChatsRepo, txManager gateways.TransactionsManager, usersApi gateways.UsersAPI) *UseCase {
	return &UseCase{
		chatsRepo: chatsRepo,
		txManager: txManager,
		usersApi:  usersApi,
	}
}

func (uc *UseCase) ListChats(ctx context.Context) ([]entity.Chat, error) {
	chats, err := uc.chatsRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (uc *UseCase) CreatePersonalChat(ctx context.Context, data *dto.CreatePersonalChat) (*entity.PersonalChat, error) {
	isExists, err := uc.usersApi.CheckExists(ctx, data.WithUserId)
	if err != nil {
		return nil, err
	}
	if !isExists {
		return nil, ErrUserNotFound
	}
	chat, err := uc.chatsRepo.Save(ctx, &entity.PersonalChat{FromUserId: data.CurrentUserId, ToUserId: data.WithUserId})
	if err != nil {
		if errors.Is(err, storage.ErrAlreadyExists) {
			return nil, ErrChatAlreadyExists
		}
		if errors.Is(err, storage.ErrRelationNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return chat.(*entity.PersonalChat), err
}

// func (uc *UseCase) CreateGroupChat(ctx context.Context, dto dto.CreateGroupChat) (*entity.PersonalChat, error)
