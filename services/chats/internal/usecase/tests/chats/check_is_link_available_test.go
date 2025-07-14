package chats_test

import (
	"context"
	"testing"

	usecase_test "github.com/AlexeyTarasov77/messanger.chats/internal/usecase/tests"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func TestCheckIsLinkAvailable(t *testing.T) {
	suite := usecase_test.NewUseCaseTestSuite(t)
	ctx := context.Background()
	fakeLinkUrl := gofakeit.UUID()
	testCases := []usecase_test.TestCase{
		{
			Name: "success",
			Mock: func() {
				suite.MockLinksRepo.EXPECT().CheckExistsByUrl(ctx, fakeLinkUrl).Return(true, nil)
			},
			Expected: true,
		},
		{
			Name: "check exists repo error",
			Mock: func() {
				suite.MockLinksRepo.EXPECT().CheckExistsByUrl(ctx, fakeLinkUrl).Return(false, fakeErr)
			},
			Expected: false,
			Err:      fakeErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Mock()
			isAvailable, err := suite.ChatsUseCase.CheckIsLinkAvailable(ctx, fakeLinkUrl)
			assert.ErrorIs(t, err, tc.Err)
			assert.Equal(t, tc.Expected, isAvailable)
		})
	}
}
