package chats_test

import (
	"context"
	"testing"

	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase/tests"
	"github.com/stretchr/testify/assert"
)

func TestListChats(t *testing.T) {
	ctx := context.Background()
	suite := usecase_test.NewUseCaseTestSuite(t)
	expectedChats := []entity.Chat{usecase_test.NewTestChat(), usecase_test.NewTestChat()}
	testCases := []usecase_test.TestCase{
		{
			Name: "repo error",
			Mock: func() {
				suite.MockChatsRepo.EXPECT().GetAll(ctx).Return(nil, fakeErr)
			},
			Expected: []entity.Chat(nil),
			Err:      fakeErr,
		},
		{
			Name: "success",
			Mock: func() {
				suite.MockChatsRepo.EXPECT().GetAll(ctx).Return(expectedChats, nil)
			},
			Expected: expectedChats,
		},
	}

	for _, tc := range testCases {
		tc.Mock()
		chats, err := suite.ChatsUseCase.ListChats(ctx)
		assert.ErrorIs(t, tc.Err, err)
		assert.Equal(t, tc.Expected.([]entity.Chat), chats)
	}
}
