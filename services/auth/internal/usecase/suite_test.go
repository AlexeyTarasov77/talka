package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/AlexeyTarasov77/messanger.users/internal/usecase/auth"
	"go.uber.org/mock/gomock"
)

type useCaseTestSuite struct {
	authUseCase               *auth.UseCase
	mockTxManager             *MockTransactionsManager
	mockUsersRepo             *MockUsersRepo
	mockOAuthProvider         *MockOAuthProvider
	mockSessionManager        *MockSessionManager
	mockSessionManagerFactory *MockSessionManagerFactory
	mockSecurityProvider      *MockSecurityProvider
	mockJwtProvider           *MockJwtProvider
	ctrl                      *gomock.Controller
}

func NewUseCaseTestSuite(t *testing.T) *useCaseTestSuite {
	ctrl := gomock.NewController(t)
	mockTxManager := NewMockTransactionsManager(ctrl)
	mockUsersRepo := NewMockUsersRepo(ctrl)
	mockSessionManager := NewMockSessionManager(ctrl)
	mockSessionManagerFactory := NewMockSessionManagerFactory(ctrl)
	mockSecurityProvider := NewMockSecurityProvider(ctrl)
	mockJwtProvider := NewMockJwtProvider(ctrl)
	chatsUseCase := auth.New(mockTxManager, mockUsersRepo, mockSessionManagerFactory, mockSecurityProvider, mockJwtProvider, time.Hour)
	return &useCaseTestSuite{
		authUseCase:               chatsUseCase,
		mockOAuthProvider:         NewMockOAuthProvider(ctrl),
		mockTxManager:             mockTxManager,
		mockUsersRepo:             mockUsersRepo,
		mockSessionManager:        mockSessionManager,
		mockSessionManagerFactory: mockSessionManagerFactory,
		mockJwtProvider:           mockJwtProvider,
		mockSecurityProvider:      mockSecurityProvider,
		ctrl:                      ctrl,
	}
}

// setTxExpectations adds expectations on mock transaction manager and returned transaction
// to ensure they are properly managed
func setTxExpectations(ctx context.Context, suite *useCaseTestSuite, commitExpected bool) {
	mockTx := NewMockTransaction(suite.ctrl)
	suite.mockTxManager.EXPECT().StartTransaction(ctx).Return(mockTx, nil)
	if commitExpected {
		commitCall := mockTx.EXPECT().Commit(ctx).Return(nil)
		// Rollback is always executed, because it's used in defer.
		// But it'll take on effect if commit was already called before
		mockTx.EXPECT().Rollback(ctx).Return(nil).After(commitCall)
		return
	}
	mockTx.EXPECT().Rollback(ctx).Return(nil)
}
