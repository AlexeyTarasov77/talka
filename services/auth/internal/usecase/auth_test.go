package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
	"github.com/stretchr/testify/assert"
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
	expectedProviders := make([]entity.OAuthProvider, 0, len(entity.OAuthSupportedProviders))
	for _, id := range entity.OAuthSupportedProviders {
		expectedProviders = append(expectedProviders, entity.OAuthProvider{ID: id, Name: id.String()})
	}
	providers := suite.authUseCase.GetOAuthProviders(ctx)

	assert.Equal(t, expectedProviders, providers)
}
