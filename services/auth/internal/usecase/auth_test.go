package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
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
	expectedOAuthToken := gofakeit.UUID()
	gomock.InOrder(
		suite.mockSecurityProvider.EXPECT().GenerateSecureUrlSafeToken().Return(expectedOAuthToken),
		suite.mockOAuthProvider.EXPECT().GetAuthURL(expectedOAuthToken).Return(expectedAuthURL),
		suite.mockSessionManagerFactory.EXPECT().CreateSessionManager(sessionId).Return(suite.mockSessionManager),
		suite.mockSessionManager.EXPECT().SetToSession(suite.authUseCase.OAuthStateTokenKey, expectedOAuthToken),
	)

	authURL, err := suite.authUseCase.SignInOAuthBegin(ctx, suite.mockOAuthProvider, sessionId)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedAuthURL, authURL)
}
