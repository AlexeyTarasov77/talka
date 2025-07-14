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

type TestCase struct {
	Name     string
	Mock     func()
	Expected any
	Err      error
}

type useCaseTestSuite struct {
	ChatsUseCase      *chats.UseCase
	MockChatsRepo     *MockChatsRepo
	MockLinksRepo     *MockInvitationLinksRepo
	MockTxManager     *MockTransactionsManager
	MockUsersRepo     *MockUsersRepo
	MockSlugGenerator *MockSlugGenerator
	ctrl              *gomock.Controller
}

func NewUseCaseTestSuite(t *testing.T) *useCaseTestSuite {
	ctrl := gomock.NewController(t)
	mockTxManager := NewMockTransactionsManager(ctrl)
	mockChatsRepo := NewMockChatsRepo(ctrl)
	mockUsersRepo := NewMockUsersRepo(ctrl)
	mockLinksRepo := NewMockInvitationLinksRepo(ctrl)
	mockSlugGenerator := NewMockSlugGenerator(ctrl)
	chatsUseCase := chats.New(mockChatsRepo, mockTxManager, mockUsersRepo, mockSlugGenerator, mockLinksRepo)
	return &useCaseTestSuite{
		ChatsUseCase:      chatsUseCase,
		MockChatsRepo:     mockChatsRepo,
		MockTxManager:     mockTxManager,
		MockUsersRepo:     mockUsersRepo,
		MockLinksRepo:     mockLinksRepo,
		MockSlugGenerator: mockSlugGenerator,
		ctrl:              ctrl,
	}
}

// SetTxExpectations adds expectations on mock transaction manager and returned transaction
// to ensure they are properly managed
func SetTxExpectations(ctx context.Context, suite *useCaseTestSuite, commitExpected bool) *MockTransaction {
	mockTx := NewMockTransaction(suite.ctrl)
	suite.MockTxManager.EXPECT().StartTransaction(ctx).Return(mockTx, nil)
	if commitExpected {
		commitCall := mockTx.EXPECT().Commit(gomock.Any()).Return(nil)
		// Rollback is always executed, because it's expected to be used in defer stmt right after starting tx.
		// But it'll take no effect if commit was already called before
		mockTx.EXPECT().Rollback(gomock.Any()).Return(nil).After(commitCall)
		return mockTx
	}
	mockTx.EXPECT().Rollback(gomock.Any()).Return(nil)
	return mockTx
}

func NewTestChat() entity.Chat {
	return &entity.BaseChat{
		ID:          gofakeit.Number(1, 999999),
		Typ:         entity.ChatTypes[rand.Intn(len(entity.ChatTypes))],
		CreatedAt:   gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
		ImageURL:    gofakeit.ImageURL(200, 200),
		LastMsgText: gofakeit.Sentence(gofakeit.Number(3, 15)),
		LastMsgDate: gofakeit.DateRange(time.Now().AddDate(0, 0, -30), time.Now()),
	}
}

func NewTestMsg() *entity.Message {
	return &entity.Message{
		ID:        gofakeit.Number(1, 999999),
		ChatId:    gofakeit.Number(1, 999999),
		Text:      gofakeit.Sentence(3),
		CreatedAt: gofakeit.DateRange(time.Now().AddDate(0, 0, -30), time.Now()),
		UpdatedAt: gofakeit.DateRange(time.Now().AddDate(0, 0, -30), time.Now()),
	}
}
