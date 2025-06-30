package chats

import (
	"context"
	"errors"

	"github.com/AlexeyTarasov77/messanger.chats/internal/dto"
	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/AlexeyTarasov77/messanger.chats/internal/gateways"
	"github.com/AlexeyTarasov77/messanger.chats/internal/gateways/storage"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase"
)

// UseCase -.
type UseCase struct {
	chatsRepo     gateways.ChatsRepo
	txManager     gateways.TransactionsManager
	usersRepo     gateways.UsersRepo
	slugGenerator gateways.SlugGenerator
}

// New -.
func New(chatsRepo gateways.ChatsRepo, txManager gateways.TransactionsManager, usersRepo gateways.UsersRepo, slugGenerator gateways.SlugGenerator) *UseCase {
	return &UseCase{
		chatsRepo:     chatsRepo,
		txManager:     txManager,
		usersRepo:     usersRepo,
		slugGenerator: slugGenerator,
	}
}

func (uc *UseCase) ListChats(ctx context.Context) ([]entity.Chat, error) {
	chats, err := uc.chatsRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (uc *UseCase) CreatePersonalChat(ctx context.Context, payload *dto.CreatePersonalChat) (*entity.PersonalChat, error) {
	isExists, err := uc.usersRepo.CheckExistsByIds(ctx, []int{payload.WithUserId})
	if err != nil {
		return nil, err
	}
	if !isExists {
		return nil, ErrUserNotFound
	}
	chat, err := uc.chatsRepo.Save(ctx, &entity.PersonalChat{FromUserId: payload.CurrentUserId, ToUserId: payload.WithUserId})
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

func (uc *UseCase) CreateGroupChat(ctx context.Context, payload *dto.CreateGroupChat) (*entity.GroupChat, error) {
	ent := payload.MapToEntity()
	if ent.Slug == "" {
		generatedSlug, err := uc.slugGenerator.GenerateRandomSlug()
		if err != nil {
			return nil, err
		}
		ent.Slug = generatedSlug
	}
	tx, err := uc.txManager.StartTransaction(ctx)
	if err != nil {
		return nil, err
	}
	ctx = usecase.SetTransaction(ctx, tx)
	defer tx.Rollback(ctx)
	chat, err := uc.chatsRepo.Save(ctx, ent)
	if err != nil {
		if errors.Is(err, storage.ErrAlreadyExists) {
			return nil, ErrChatAlreadyExists
		}
		return nil, err
	}
	if len(payload.MembersIds) > 0 {
		isExists, err := uc.usersRepo.CheckExistsByIds(ctx, payload.MembersIds)
		if err != nil {
			return nil, err
		}
		if !isExists {
			return nil, ErrMemberNotFound
		}
		if err = uc.chatsRepo.AddMembers(ctx, chat.GetID(), payload.MembersIds); err != nil {
			return nil, ErrMemberAlreadyInChat
		}
	}
	tx.Commit(ctx)
	return chat.(*entity.GroupChat), nil
}

func (uc *UseCase) UpdateLastMsg(ctx context.Context, msg *entity.Message) error {
	err := uc.chatsRepo.UpdateLastMsgInfo(ctx, msg.ChatId, msg.Text, msg.CreatedAt)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return ErrChatNotFound
		}
		return err
	}
	return nil
}
