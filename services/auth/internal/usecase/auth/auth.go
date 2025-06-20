package auth

import (
	"context"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways"
)

// UseCase -.
type UseCase struct {
	usersRepo             gateways.UsersRepo
	txManager             gateways.TransactionsManager
	sessionManagerFactory gateways.SessionManagerFactory
	securityProvider      gateways.SecurityProvider
	OAuthStateTokenKey    string
}

// New -.
func New(txManager gateways.TransactionsManager, usersRepo gateways.UsersRepo, sessionManagerFactory gateways.SessionManagerFactory, securityProvider gateways.SecurityProvider) *UseCase {
	return &UseCase{
		txManager:             txManager,
		usersRepo:             usersRepo,
		sessionManagerFactory: sessionManagerFactory,
		securityProvider:      securityProvider,
		OAuthStateTokenKey:    "oauth_state_token",
	}
}

func (uc *UseCase) SignIn(ctx context.Context) (*entity.User, error) {
	return nil, nil
}

func (uc *UseCase) GetOAuthProviders(ctx context.Context) []entity.OAuthProviderInfo {
	providers := make([]entity.OAuthProviderInfo, 0, len(entity.OAuthSupportedProviders))
	for _, id := range entity.OAuthSupportedProviders {
		providers = append(providers, entity.OAuthProviderInfo{ID: id, Name: id.String()})
	}
	return providers
}

func (uc *UseCase) SignInOAuthBegin(ctx context.Context, provider gateways.OAuthProvider, sessionId string) (string, error) {
	stateToken := uc.securityProvider.GenerateSecureUrlSafeToken()
	authURL := provider.GetAuthURL(stateToken)
	sessionManager := uc.sessionManagerFactory.CreateSessionManager(sessionId)
	if err := sessionManager.SetToSession(uc.OAuthStateTokenKey, stateToken); err != nil {
		return "", err
	}
	return authURL, nil
}

func (uc *UseCase) SignInOAuthComplete(ctx context.Context, provider gateways.OAuthProvider) (*entity.User, error) {
	return nil, nil
}

func (uc *UseCase) SignUp(ctx context.Context) (*entity.User, error) {
	return nil, nil
}
