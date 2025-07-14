package chats_test

import (
	"context"
	"testing"

	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/AlexeyTarasov77/messanger.chats/internal/gateways/storage"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase/chats"
	usecase_test "github.com/AlexeyTarasov77/messanger.chats/internal/usecase/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetChat(t *testing.T) {
	suite := usecase_test.NewUseCaseTestSuite(t)
	ctx := context.Background()
	fakeChat := usecase_test.NewTestChat()
	fakeMessages := []entity.Message{*usecase_test.NewTestMsg(), *usecase_test.NewTestMsg()}
	expectedChat := fakeChat.(*entity.BaseChat)
	expectedChat.Messages = fakeMessages
	testCases := []usecase_test.TestCase{
		{
			Name: "success",
			Mock: func() {
				suite.MockChatsRepo.EXPECT().GetById(ctx, fakeChat.GetID()).Return(fakeChat, nil)
				suite.MockMessagesRepo.EXPECT().GetByChatId(ctx, fakeChat.GetID()).Return(fakeMessages, nil)
			},
			Expected: expectedChat,
		},
		{
			Name: "not-found",
			Mock: func() {
				suite.MockChatsRepo.EXPECT().GetById(ctx, fakeChat.GetID()).Return(nil, storage.ErrNotFound)
			},
			Err: chats.ErrChatNotFound,
		},
		{
			Name: "chats-repo-get-by-id-error",
			Mock: func() {
				suite.MockChatsRepo.EXPECT().GetById(ctx, fakeChat.GetID()).Return(nil, fakeErr)
			},
			Err: fakeErr,
		},
		{
			Name: "messages-repo-get-by-chatid-error",
			Mock: func() {
				suite.MockChatsRepo.EXPECT().GetById(ctx, fakeChat.GetID()).Return(fakeChat, nil)
				suite.MockMessagesRepo.EXPECT().GetByChatId(ctx, fakeChat.GetID()).Return(nil, fakeErr)
			},
			Err: fakeErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Mock()
			chat, err := suite.ChatsUseCase.GetChat(ctx, fakeChat.GetID())
			assert.ErrorIs(t, err, tc.Err)
			assert.Equal(t, tc.Expected, chat)
		})
	}
}
