package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/storage"
)

// UseCase -.
type UseCase struct {
	usersRepo             gateways.UsersRepo
	txManager             gateways.TransactionsManager
	sessionManagerFactory gateways.SessionManagerFactory
	securityProvider      gateways.SecurityProvider
	OAuthStateTokenKey    string
	AuthTokenTTL          time.Duration
	jwtProvider           gateways.JwtProvider
}

// New -.
func New(txManager gateways.TransactionsManager, usersRepo gateways.UsersRepo, sessionManagerFactory gateways.SessionManagerFactory, securityProvider gateways.SecurityProvider, jwtProvider gateways.JwtProvider, authTokenTTL time.Duration) *UseCase {
	return &UseCase{
		txManager:             txManager,
		usersRepo:             usersRepo,
		sessionManagerFactory: sessionManagerFactory,
		securityProvider:      securityProvider,
		AuthTokenTTL:          authTokenTTL,
		OAuthStateTokenKey:    "oauth_state_token",
		jwtProvider:           jwtProvider,
	}
}

type AuthInfo struct {
	Token    string
	User     *entity.User
	SignedUp bool
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
	if err := sessionManager.SetToSession(ctx, uc.OAuthStateTokenKey, stateToken); err != nil {
		return "", err
	}
	return authURL, nil
}

func (uc *UseCase) SignInOAuthComplete(ctx context.Context, stateToken string, authCode string, provider gateways.OAuthProvider, sessionId string) (*AuthInfo, error) {
	var signedUp bool
	sessionManager := uc.sessionManagerFactory.CreateSessionManager(sessionId)
	sessionData, err := sessionManager.GetSessionData(ctx)
	if err != nil {
		return nil, err
	}
	stateTokenFromSession, ok := sessionData[uc.OAuthStateTokenKey]
	if !ok {
		return nil, ErrOAuthSignInNotStarted
	}
	if stateTokenFromSession != stateToken {
		return nil, fmt.Errorf("invalid state token: %w", ErrOAuthSignInFailed)
	}
	oauthAccessToken, err := provider.GetAccessToken(ctx, authCode)
	userData, err := provider.FetchUserData(ctx, oauthAccessToken)
	if err != nil {
		return nil, err
	}
	user, err := uc.usersRepo.GetByOAuthAccId(ctx, userData.OAuthAccID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			user, err = uc.usersRepo.Insert(ctx, userData)
			signedUp = true
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	token, err := uc.jwtProvider.NewToken(uc.AuthTokenTTL, map[string]any{"uid": user.ID})
	if err != nil {
		return nil, err
	}
	return &AuthInfo{User: user, Token: token, SignedUp: signedUp}, nil
}

func (uc *UseCase) SignUp(ctx context.Context) (*entity.User, error) {
	return nil, nil
}
