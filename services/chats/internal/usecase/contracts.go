// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/AlexeyTarasov77/messanger.chats/internal/dto"
	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

type (
	Chats interface {
		CreatePersonalChat(ctx context.Context, dto dto.CreatePersonalChat) (*entity.PersonalChat, error)
		CreateGroupChat(ctx context.Context, dto dto.CreateGroupChat) (*entity.PersonalChat, error)
	}
)
