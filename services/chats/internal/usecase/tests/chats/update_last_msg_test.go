package chats_test

import (
	"context"
	"testing"

	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/AlexeyTarasov77/messanger.chats/internal/gateways/storage"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase/chats"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase/tests"
	"github.com/stretchr/testify/assert"
)

func TestUpdateLastMsg(t *testing.T) {
	suite := usecase_test.NewUseCaseTestSuite(t)
	ctx := context.Background()
	var expectedMsg *entity.Message
	testCases := []usecase_test.TestCase{
		{
			Name: "success",
			Mock: func() {
				suite.MockChatsRepo.EXPECT().Update(ctx, expectedMsg.ChatId, map[string]any{"last_msg": expectedMsg})
			},
		},
		{
			Name: "update error (unhandled)",
			Mock: func() {
				suite.MockChatsRepo.EXPECT().Update(ctx, expectedMsg.ChatId, map[string]any{"last_msg": expectedMsg}).Return(fakeErr)
			},
			Err: fakeErr,
		},
		{
			Name: "update error (not found)",
			Mock: func() {
				suite.MockChatsRepo.EXPECT().Update(ctx, expectedMsg.ChatId, map[string]any{"last_msg": expectedMsg}).Return(storage.ErrNotFound)
			},
			Err: chats.ErrChatNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			expectedMsg = usecase_test.NewTestMsg()
			tc.Mock()

			err := suite.ChatsUseCase.UpdateLastMsg(ctx, expectedMsg)
			assert.ErrorIs(t, err, tc.Err)
		})
	}
}
