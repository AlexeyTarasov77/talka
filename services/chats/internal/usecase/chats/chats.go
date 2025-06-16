package chats

import (
	"context"

	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/AlexeyTarasov77/messanger.chats/internal/gateways"
)

// UseCase -.
type UseCase struct {
	chatsRepo gateways.ChatsRepo
	txManager gateways.TransactionsManager
}

// New -.
func New(chatsRepo gateways.ChatsRepo, txManager gateways.TransactionsManager) *UseCase {
	return &UseCase{
		chatsRepo: chatsRepo,
		txManager: txManager,
	}
}

func (uc *UseCase) ListChats(ctx context.Context) ([]entity.Chat, error) {
	chats, err := uc.chatsRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

// func (uc *UseCase) CreatePersonalChat(ctx context.Context, dto dto.CreatePersonalChat) (*entity.PersonalChat, error)
// func (uc *UseCase) CreateGroupChat(ctx context.Context, dto dto.CreateGroupChat) (*entity.PersonalChat, error)
