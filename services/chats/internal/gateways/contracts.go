// Package gateways implements application outer layer logic. Each logic group in own file.
package gateways

import (
	"context"
	"time"

	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
)

//go:generate mockgen -source=contracts.go -destination=../usecase/tests/mocks_gateways.go -package=usecase_test

type (
	ChatsRepo interface {
		GetAll(ctx context.Context) ([]entity.Chat, error)
		Save(ctx context.Context, chat entity.Chat) (entity.Chat, error)
		UpdateLastMsgInfo(ctx context.Context, chatId int, msgText string, msgDate time.Time) error
		AddMembers(ctx context.Context, groupChatId int, membersIds []int) error
		GetGroupByLink(ctx context.Context, link string) (*entity.GroupChat, error)
		CreateJoinReq(ctx context.Context, userId, chatId int) error
		CountMembersByLink(ctx context.Context, groupId int, linkId int) (int, error)
		CountJoinRequestsByLink(ctx context.Context, groupId int, linkId int) (int, error)
	}
	InvitationLinksRepo interface {
		CheckExistsByUrl(ctx context.Context, url string) (bool, error)
		GetByUrl(ctx context.Context, url string) (*entity.InvitationLink, error)
	}
	UsersRepo interface {
		CheckExistsByIds(ctx context.Context, ids []int) (bool, error)
	}
	Transaction interface {
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}
	TransactionsManager interface {
		StartTransaction(ctx context.Context) (Transaction, error)
	}
	SlugGenerator interface {
		GenerateRandomSlug() (string, error)
	}
)
