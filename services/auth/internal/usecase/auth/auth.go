package auth

import (
	"context"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways"
	oauth_providers "github.com/AlexeyTarasov77/messanger.users/internal/gateways/oauth/providers"
)

// UseCase -.
type UseCase struct {
	usersRepo gateways.UsersRepo
	txManager gateways.TransactionsManager
}

// New -.
func New(txManager gateways.TransactionsManager, usersRepo gateways.UsersRepo) *UseCase {
	return &UseCase{
		txManager: txManager,
		usersRepo: usersRepo,
	}
}

func (uc *UseCase) SignIn(ctx context.Context) (*entity.User, error) {
	return nil, nil
}

func (uc *UseCase) GetOAuthProviders(ctx context.Context) []entity.OAuthProvider {
	providers := make([]entity.OAuthProvider, 0, len(entity.OAuthSupportedProviders))
	for _, id := range entity.OAuthSupportedProviders {
		providers = append(providers, entity.OAuthProvider{ID: id, Name: id.String()})
	}
	return providers
}

func (uc *UseCase) SignInOAuth(ctx context.Context, provider oauth_providers.Interface) (*entity.User, error) {
	return nil, nil
}

func (uc *UseCase) SignUp(ctx context.Context) (*entity.User, error) {
	return nil, nil
}
