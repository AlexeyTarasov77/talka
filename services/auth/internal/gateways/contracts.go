// Package gateways implements application outer layer logic. Each logic group in own file.
package gateways

import (
	"context"
	"time"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
)

//go:generate mockgen -source=contracts.go -destination=../usecase/mocks_gateways_test.go -package=usecase_test

type (
	UsersRepo interface {
		Insert(ctx context.Context, user *entity.User) (*entity.User, error)
		GetByOAuthAccId(ctx context.Context, accId string) (*entity.User, error)
	}
	Transaction interface {
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}
	TransactionsManager interface {
		StartTransaction(ctx context.Context) (Transaction, error)
	}
	OAuthProvider interface {
		GetAuthURL(stateToken string) string
		GetAccessToken(ctx context.Context, authCode string) (string, error)
		FetchUserData(ctx context.Context, accessToken string) (*entity.User, error)
	}
	SessionManager interface {
		GetSessionData(ctx context.Context) (map[string]any, error)
		SetToSession(ctx context.Context, key, value string) error
	}
	SessionManagerFactory interface {
		CreateSessionManager(sessionId string) SessionManager
	}
	JwtProvider interface {
		NewToken(expires time.Duration, claims map[string]any) (string, error)
	}
	SecurityProvider interface {
		GenerateSecureUrlSafeToken() string
	}
)
