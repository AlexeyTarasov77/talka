package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/AlexeyTarasov77/messanger.chats/internal/dto"
	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/AlexeyTarasov77/messanger.chats/internal/gateways/storage"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase/chats"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

var fakeErr = errors.New("fake error")

type testCase struct {
	name     string
	mock     func()
	expected any
	err      error
}

func TestListChats(t *testing.T) {
	ctx := context.Background()
	suite := NewUseCaseTestSuite(t)
	expectedChats := []entity.Chat{newTestChat(), newTestChat()}
	testCases := []testCase{
		{
			name: "repo error",
			mock: func() {
				suite.mockChatsRepo.EXPECT().GetAll(ctx).Return(nil, fakeErr)
			},
			expected: []entity.Chat(nil),
			err:      fakeErr,
		},
		{
			name: "success",
			mock: func() {
				suite.mockChatsRepo.EXPECT().GetAll(ctx).Return(expectedChats, nil)
			},
			expected: expectedChats,
		},
	}

	for _, tc := range testCases {
		tc.mock()
		chats, err := suite.chatsUseCase.ListChats(ctx)
		assert.ErrorIs(t, tc.err, err)
		assert.Equal(t, tc.expected.([]entity.Chat), chats)
	}
}

func TestCreatePersonalChat(t *testing.T) {
	ctx := context.Background()
	suite := NewUseCaseTestSuite(t)
	payload := &dto.CreatePersonalChat{CurrentUserId: gofakeit.Number(1, 999999), WithUserId: gofakeit.Number(1, 999999)}
	chatToInsert := &entity.PersonalChat{FromUserId: payload.CurrentUserId, ToUserId: payload.WithUserId}
	expectedChat := &entity.PersonalChat{Chat: newTestChat()}
	testCases := []testCase{
		{
			name: "save error (unhandled)",
			mock: func() {
				checkExistsCall := suite.mockUsersApi.EXPECT().CheckExists(ctx, payload.WithUserId).Return(true, nil)
				suite.mockChatsRepo.EXPECT().Save(ctx, chatToInsert).Return(nil, fakeErr).After(checkExistsCall)
			},
			expected: (*entity.PersonalChat)(nil),
			err:      fakeErr,
		},
		{
			name: "save error - already exists",
			mock: func() {
				checkExistsCall := suite.mockUsersApi.EXPECT().CheckExists(ctx, payload.WithUserId).Return(true, nil)
				suite.mockChatsRepo.EXPECT().Save(ctx, chatToInsert).Return(nil, storage.ErrAlreadyExists).After(checkExistsCall)
			},
			expected: (*entity.PersonalChat)(nil),
			err:      chats.ErrChatAlreadyExists,
		},
		{
			name: "user not found",
			mock: func() {
				suite.mockUsersApi.EXPECT().CheckExists(ctx, payload.WithUserId).Return(false, nil)
			},
			expected: (*entity.PersonalChat)(nil),
			err:      chats.ErrUserNotFound,
		},
		{
			name: "usersApi.checkExists error (unhandled)",
			mock: func() {
				suite.mockUsersApi.EXPECT().CheckExists(ctx, payload.WithUserId).Return(false, fakeErr)
			},
			expected: (*entity.PersonalChat)(nil),
			err:      fakeErr,
		},
		{
			name: "success",
			mock: func() {
				checkExistsCall := suite.mockUsersApi.EXPECT().CheckExists(ctx, payload.WithUserId).Return(true, nil)
				suite.mockChatsRepo.EXPECT().Save(ctx, chatToInsert).Return(expectedChat, nil).After(checkExistsCall)
			},
			expected: expectedChat,
		},
	}

	for _, tc := range testCases {
		tc.mock()
		chat, err := suite.chatsUseCase.CreatePersonalChat(ctx, payload)
		assert.ErrorIs(t, err, tc.err)
		assert.Equal(t, tc.expected, chat)
	}
}
