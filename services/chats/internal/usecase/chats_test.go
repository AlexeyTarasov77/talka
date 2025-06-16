package usecase_test

import (
	"context"
	"testing"

	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestListChats(t *testing.T) {
	ctx := context.Background()
	suite := NewUseCaseTestSuite(t)
	expectedChats := []entity.Chat{newTestChat(), newTestChat()}
	suite.mockChatsRepo.EXPECT().GetAll(ctx).Return(expectedChats, nil)

	chats, err := suite.chatsUseCase.ListChats(ctx)

	assert.NoError(t, err)
	for i, chat := range chats {
		assert.Equal(t, expectedChats[i].GetID(), chat.GetID())
	}
}
