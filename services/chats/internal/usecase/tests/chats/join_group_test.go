package chats_test

import (
	"context"
	"testing"
	"time"

	"github.com/AlexeyTarasov77/messanger.chats/internal/entity"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase/chats"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase/tests"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func TestJoinGroupChat(t *testing.T) {
	suite := usecase_test.NewUseCaseTestSuite(t)
	ctx := context.Background()
	fakeUserId := gofakeit.Number(1, 100)
	fakeLinkUrl := gofakeit.UUID()
	testCases := []usecase_test.TestCase{
		{
			Name: "with primary link",
			Mock: func() {
				tx := usecase_test.SetTxExpectations(ctx, suite, true)
				txCtx := usecase.SetTransaction(ctx, tx)
				fakeGroupChat := &entity.GroupChat{Chat: usecase_test.NewTestChat(), PrimaryLinkUrl: fakeLinkUrl}
				suite.MockChatsRepo.EXPECT().GetGroupByLink(txCtx, fakeLinkUrl).Return(fakeGroupChat, nil)
				suite.MockChatsRepo.EXPECT().CreateJoinReq(txCtx, fakeUserId, fakeGroupChat.GetID()).Return(nil)
			},
			Expected: false,
		},
		{
			Name: "immediate join",
			Mock: func() {
				tx := usecase_test.SetTxExpectations(ctx, suite, true)
				txCtx := usecase.SetTransaction(ctx, tx)
				fakeGroupChat := &entity.GroupChat{Chat: usecase_test.NewTestChat()}
				fakeLink := &entity.InvitationLink{Url: fakeLinkUrl, GroupId: fakeGroupChat.GetID()}
				suite.MockChatsRepo.EXPECT().GetGroupByLink(txCtx, fakeLinkUrl).Return(fakeGroupChat, nil)
				suite.MockLinksRepo.EXPECT().GetByUrl(txCtx, fakeLinkUrl).Return(fakeLink, nil)
				suite.MockChatsRepo.EXPECT().AddMembers(txCtx, fakeGroupChat.GetID(), []int{fakeUserId}).Return(nil)
			},
			Expected: true,
		},
		{
			Name: "link requires admin approval",
			Mock: func() {
				tx := usecase_test.SetTxExpectations(ctx, suite, true)
				txCtx := usecase.SetTransaction(ctx, tx)
				fakeGroupChat := &entity.GroupChat{Chat: usecase_test.NewTestChat()}
				fakeLink := &entity.InvitationLink{Url: fakeLinkUrl, GroupId: fakeGroupChat.GetID(), RequiresAdminApproval: true}
				suite.MockChatsRepo.EXPECT().GetGroupByLink(txCtx, fakeLinkUrl).Return(fakeGroupChat, nil)
				suite.MockLinksRepo.EXPECT().GetByUrl(txCtx, fakeLinkUrl).Return(fakeLink, nil)
				suite.MockChatsRepo.EXPECT().CreateJoinReq(txCtx, fakeUserId, fakeGroupChat.GetID()).Return(nil)
			},
			Expected: false,
		},
		// TODO: if link requires admin approval also count amount of join requests created by link
		{
			Name: "link activations limit exceeded (members)",
			Mock: func() {
				tx := usecase_test.SetTxExpectations(ctx, suite, false)
				txCtx := usecase.SetTransaction(ctx, tx)
				fakeGroupChat := &entity.GroupChat{Chat: usecase_test.NewTestChat()}
				fakeLink := &entity.InvitationLink{Url: fakeLinkUrl, GroupId: fakeGroupChat.GetID(), ActivationsLimit: 1}
				suite.MockChatsRepo.EXPECT().GetGroupByLink(txCtx, fakeLinkUrl).Return(fakeGroupChat, nil)
				suite.MockLinksRepo.EXPECT().GetByUrl(txCtx, fakeLinkUrl).Return(fakeLink, nil)
				suite.MockChatsRepo.EXPECT().CountMembersByLink(txCtx, fakeGroupChat.GetID(), fakeLink.ID).Return(gofakeit.Number(1, 999), nil)
			},
			Err:      chats.ErrLinkActivationsLimitExceeded,
			Expected: false,
		},
		{
			Name: "link activations limit exceeded (members+join-reqs)",
			Mock: func() {
				tx := usecase_test.SetTxExpectations(ctx, suite, false)
				txCtx := usecase.SetTransaction(ctx, tx)
				fakeGroupChat := &entity.GroupChat{Chat: usecase_test.NewTestChat()}
				fakeLink := &entity.InvitationLink{Url: fakeLinkUrl, GroupId: fakeGroupChat.GetID(), ActivationsLimit: 10, RequiresAdminApproval: true}
				suite.MockChatsRepo.EXPECT().GetGroupByLink(txCtx, fakeLinkUrl).Return(fakeGroupChat, nil)
				suite.MockLinksRepo.EXPECT().GetByUrl(txCtx, fakeLinkUrl).Return(fakeLink, nil)
				// use exact range to ensure that both return values are summed up for check
				suite.MockChatsRepo.EXPECT().CountMembersByLink(txCtx, fakeGroupChat.GetID(), fakeLink.ID).Return(gofakeit.Number(5, 9), nil)
				suite.MockChatsRepo.EXPECT().CountJoinRequestsByLink(txCtx, fakeGroupChat.GetID(), fakeLink.ID).Return(gofakeit.Number(5, 9), nil)
			},
			Err:      chats.ErrLinkActivationsLimitExceeded,
			Expected: false,
		},
		{
			Name: "link expired",
			Mock: func() {
				tx := usecase_test.SetTxExpectations(ctx, suite, false)
				txCtx := usecase.SetTransaction(ctx, tx)
				fakeGroupChat := &entity.GroupChat{Chat: usecase_test.NewTestChat()}
				fakeLink := &entity.InvitationLink{Url: fakeLinkUrl, GroupId: fakeGroupChat.GetID(), ExpiresAt: time.Now()}
				suite.MockChatsRepo.EXPECT().GetGroupByLink(txCtx, fakeLinkUrl).Return(fakeGroupChat, nil)
				suite.MockLinksRepo.EXPECT().GetByUrl(txCtx, fakeLinkUrl).Return(fakeLink, nil)
			},
			Err:      chats.ErrLinkExpired,
			Expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Mock()
			isJoined, err := suite.ChatsUseCase.JoinGroupChat(ctx, fakeUserId, fakeLinkUrl)
			assert.ErrorIs(t, err, tc.Err)
			assert.Equal(t, tc.Expected, isJoined)
		})
	}
}
