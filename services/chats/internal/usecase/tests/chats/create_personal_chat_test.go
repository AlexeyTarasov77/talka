package chats_test

import (
	"context"
	"testing"

	"github.com/AlexeyTarasov77/messanger.chats/internal/dto"
	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/AlexeyTarasov77/messanger.chats/internal/gateways/storage"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase/chats"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase/tests"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func TestCreatePersonalChat(t *testing.T) {
	ctx := context.Background()
	suite := usecase_test.NewUseCaseTestSuite(t)
	payload := &dto.CreatePersonalChat{CurrentUserId: gofakeit.Number(1, 999999), WithUserId: gofakeit.Number(1, 999999)}
	chatToInsert := &entity.PersonalChat{FromUserId: payload.CurrentUserId, ToUserId: payload.WithUserId}
	expectedChat := &entity.PersonalChat{Chat: usecase_test.NewTestChat()}
	testCases := []usecase_test.TestCase{
		{
			Name: "save error (unhandled)",
			Mock: func() {
				checkExistsCall := suite.MockUsersRepo.EXPECT().CheckExistsByIds(ctx, []int{payload.WithUserId}).Return(true, nil)
				suite.MockChatsRepo.EXPECT().Save(ctx, chatToInsert).Return(nil, fakeErr).After(checkExistsCall)
			},
			Expected: (*entity.PersonalChat)(nil),
			Err:      fakeErr,
		},
		{
			Name: "save error - already exists",
			Mock: func() {
				checkExistsCall := suite.MockUsersRepo.EXPECT().CheckExistsByIds(ctx, []int{payload.WithUserId}).Return(true, nil)
				suite.MockChatsRepo.EXPECT().Save(ctx, chatToInsert).Return(nil, storage.ErrAlreadyExists).After(checkExistsCall)
			},
			Expected: (*entity.PersonalChat)(nil),
			Err:      chats.ErrChatAlreadyExists,
		},
		{
			Name: "user not found",
			Mock: func() {
				suite.MockUsersRepo.EXPECT().CheckExistsByIds(ctx, []int{payload.WithUserId}).Return(false, nil)
			},
			Expected: (*entity.PersonalChat)(nil),
			Err:      chats.ErrUserNotFound,
		},
		{
			Name: "usersRepo.checkExistsAll error (unhandled)",
			Mock: func() {
				suite.MockUsersRepo.EXPECT().CheckExistsByIds(ctx, []int{payload.WithUserId}).Return(false, fakeErr)
			},
			Expected: (*entity.PersonalChat)(nil),
			Err:      fakeErr,
		},
		{
			Name: "success",
			Mock: func() {
				checkExistsCall := suite.MockUsersRepo.EXPECT().CheckExistsByIds(ctx, []int{payload.WithUserId}).Return(true, nil)
				suite.MockChatsRepo.EXPECT().Save(ctx, chatToInsert).Return(expectedChat, nil).After(checkExistsCall)
			},
			Expected: expectedChat,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Mock()
			chat, err := suite.ChatsUseCase.CreatePersonalChat(ctx, payload)
			assert.ErrorIs(t, err, tc.Err)
			assert.Equal(t, tc.Expected, chat)
		})
	}
}
