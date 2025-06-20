package usecase_test

import (
	"context"
	"testing"

	"github.com/AlexeyTarasov77/messanger.users/internal/usecase/auth"
	"go.uber.org/mock/gomock"
)

type useCaseTestSuite struct {
	authUseCase   *auth.UseCase
	mockTxManager *MockTransactionsManager
	mockUsersRepo *MockUsersRepo
	ctrl          *gomock.Controller
}

func NewUseCaseTestSuite(t *testing.T) *useCaseTestSuite {
	ctrl := gomock.NewController(t)
	mockTxManager := NewMockTransactionsManager(ctrl)
	mockUsersRepo := NewMockUsersRepo(ctrl)
	chatsUseCase := auth.New(mockTxManager, mockUsersRepo)
	return &useCaseTestSuite{
		authUseCase:   chatsUseCase,
		mockTxManager: mockTxManager,
		mockUsersRepo: mockUsersRepo,
		ctrl:          ctrl,
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
