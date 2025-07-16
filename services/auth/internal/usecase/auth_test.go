package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/storage"
	"github.com/AlexeyTarasov77/messanger.users/internal/usecase/auth"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var fakeErr = errors.New("fake error")

type testCase struct {
	name     string
	mock     func()
	expected any
	err      error
}

func TestGetOAuthProviders(t *testing.T) {
	suite := NewUseCaseTestSuite(t)
	ctx := context.Background()
	expectedProviders := make([]entity.OAuthProviderInfo, 0, len(entity.OAuthSupportedProviders))
	for _, id := range entity.OAuthSupportedProviders {
		expectedProviders = append(expectedProviders, entity.OAuthProviderInfo{ID: id, Name: id.String()})
	}
	providers := suite.authUseCase.GetOAuthProviders(ctx)

	assert.Equal(t, expectedProviders, providers)
}

func TestSignInOAuthBegin(t *testing.T) {
	suite := NewUseCaseTestSuite(t)
	ctx := context.Background()
	sessionId := gofakeit.UUID()
	expectedAuthURL := gofakeit.URL()
	expectedOAuthStateToken := gofakeit.UUID()
	gomock.InOrder(
		suite.mockSecurityProvider.EXPECT().GenerateSecureUrlSafeToken().Return(expectedOAuthStateToken),
		suite.mockOAuthProvider.EXPECT().GetAuthURL(expectedOAuthStateToken).Return(expectedAuthURL),
		suite.mockSessionManagerFactory.EXPECT().CreateSessionManager(sessionId).Return(suite.mockSessionManager),
		suite.mockSessionManager.EXPECT().SetToSession(ctx, suite.authUseCase.OAuthStateTokenKey, expectedOAuthStateToken),
	)

	authURL, err := suite.authUseCase.SignInOAuthBegin(ctx, suite.mockOAuthProvider, sessionId)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedAuthURL, authURL)
}

func TestSignInOAuthComplete(t *testing.T) {
	suite := NewUseCaseTestSuite(t)
	ctx := context.Background()
	sessionId := gofakeit.UUID()
	expectedOAuthStateToken := gofakeit.UUID()
	oauthProvidedUserData := &entity.User{OAuthAccID: gofakeit.UUID(), Email: gofakeit.Email(), Username: gofakeit.Username()}
	expectedUser := &entity.User{ID: gofakeit.Number(1, 100)}
	fakeAuthCode := gofakeit.UUID()
	expectedAuthToken := gofakeit.UUID()
	expectedOAuthAccessToken := gofakeit.UUID()
	testCases := []testCase{
		{
			name: "success",
			mock: func() {
				gomock.InOrder(
					suite.mockSessionManagerFactory.EXPECT().CreateSessionManager(sessionId).Return(suite.mockSessionManager),
					suite.mockSessionManager.EXPECT().GetSessionData(ctx).
						Return(map[string]any{suite.authUseCase.OAuthStateTokenKey: expectedOAuthStateToken}, nil),
					suite.mockOAuthProvider.EXPECT().GetAccessToken(ctx, fakeAuthCode).Return(expectedOAuthAccessToken, nil),
					suite.mockOAuthProvider.EXPECT().FetchUserData(ctx, expectedOAuthAccessToken).Return(oauthProvidedUserData, nil),
					suite.mockUsersRepo.EXPECT().GetByOAuthAccId(ctx, oauthProvidedUserData.OAuthAccID).Return(expectedUser, nil),
					suite.mockJwtProvider.EXPECT().NewToken(suite.authUseCase.AuthTokenTTL, map[string]any{"uid": expectedUser.ID}).
						Return(expectedAuthToken, nil),
				)
			},
			expected: &auth.AuthInfo{User: expectedUser, Token: expectedAuthToken},
		},
		{
			name: "success - create new user",
			mock: func() {
				gomock.InOrder(
					suite.mockSessionManagerFactory.EXPECT().CreateSessionManager(sessionId).Return(suite.mockSessionManager),
					suite.mockSessionManager.EXPECT().GetSessionData(ctx).
						Return(map[string]any{suite.authUseCase.OAuthStateTokenKey: expectedOAuthStateToken}, nil),
					suite.mockOAuthProvider.EXPECT().GetAccessToken(ctx, fakeAuthCode).Return(expectedOAuthAccessToken, nil),
					suite.mockOAuthProvider.EXPECT().FetchUserData(ctx, expectedOAuthAccessToken).Return(oauthProvidedUserData, nil),
					suite.mockUsersRepo.EXPECT().GetByOAuthAccId(ctx, oauthProvidedUserData.OAuthAccID).Return(nil, storage.ErrNotFound),
					suite.mockUsersRepo.EXPECT().Insert(ctx, oauthProvidedUserData).Return(expectedUser, nil),
					suite.mockJwtProvider.EXPECT().NewToken(suite.authUseCase.AuthTokenTTL, map[string]any{"uid": expectedUser.ID}).
						Return(expectedAuthToken, nil),
				)
			},
			expected: &auth.AuthInfo{User: expectedUser, Token: expectedAuthToken, SignedUp: true},
		},
		{
			name: "invalid state token",
			mock: func() {
				gomock.InOrder(
					suite.mockSessionManagerFactory.EXPECT().CreateSessionManager(sessionId).Return(suite.mockSessionManager),
					suite.mockSessionManager.EXPECT().GetSessionData(ctx).
						Return(map[string]any{suite.authUseCase.OAuthStateTokenKey: "unexpected"}, nil),
				)
			},
			err:      auth.ErrOAuthSignInFailed,
			expected: (*auth.AuthInfo)(nil),
		},
		{
			name: "missing state token in session",
			mock: func() {
				gomock.InOrder(
					suite.mockSessionManagerFactory.EXPECT().CreateSessionManager(sessionId).Return(suite.mockSessionManager),
					suite.mockSessionManager.EXPECT().GetSessionData(ctx).
						Return(map[string]any{}, nil),
				)
			},
			err:      auth.ErrOAuthSignInNotStarted,
			expected: (*auth.AuthInfo)(nil),
		},
		{
			name: "failed to get session data",
			mock: func() {
				gomock.InOrder(
					suite.mockSessionManagerFactory.EXPECT().CreateSessionManager(sessionId).Return(suite.mockSessionManager),
					suite.mockSessionManager.EXPECT().GetSessionData(ctx).
						Return(nil, fakeErr),
				)
			},
			err:      fakeErr,
			expected: (*auth.AuthInfo)(nil),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			authInfo, err := suite.authUseCase.SignInOAuthComplete(
				ctx, expectedOAuthStateToken, fakeAuthCode, suite.mockOAuthProvider, sessionId,
			)
			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, tc.expected.(*auth.AuthInfo), authInfo)
		})
	}
}
