package usecase_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase/chats"
	"github.com/brianvoe/gofakeit"
	"go.uber.org/mock/gomock"
)

type useCaseTestSuite struct {
	chatsUseCase  *chats.UseCase
	mockChatsRepo *MockChatsRepo
	mockTxManager *MockTransactionsManager
	ctrl          *gomock.Controller
}

func NewUseCaseTestSuite(t *testing.T) *useCaseTestSuite {
	ctrl := gomock.NewController(t)
	mockTxManager := NewMockTransactionsManager(ctrl)
	mockChatsRepo := NewMockChatsRepo(ctrl)
	chatsUseCase := chats.New(mockChatsRepo, mockTxManager)
	return &useCaseTestSuite{
		chatsUseCase:  chatsUseCase,
		mockChatsRepo: mockChatsRepo,
		mockTxManager: mockTxManager,
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

func newTestChat() entity.Chat {
	return &entity.BaseChat{
		ID:          gofakeit.Number(1, 999999),
		Typ:         entity.ChatTypes[rand.Intn(len(entity.ChatTypes))],
		CreatedAt:   gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
		ImageURL:    gofakeit.ImageURL(200, 200),
		LastMsgText: gofakeit.Sentence(gofakeit.Number(3, 15)),
		LastMsgDate: gofakeit.DateRange(time.Now().AddDate(0, 0, -30), time.Now()),
	}
}
