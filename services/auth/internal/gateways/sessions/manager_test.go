package sessions_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/sessions"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var fakeErr = errors.New("fake error")

type testCase struct {
	name     string
	mock     func()
	expected any
	err      error
}

type testSuite struct {
	mockStorage    *MockStorage
	sessionManager *sessions.Manager
	sessionId      string
}

func NewTestSuite(t *testing.T) *testSuite {
	ctrl := gomock.NewController(t)
	mockStorage := NewMockStorage(ctrl)
	sessionId := gofakeit.UUID()
	return &testSuite{
		mockStorage:    mockStorage,
		sessionId:      sessionId,
		sessionManager: sessions.NewManager(sessionId, mockStorage),
	}
}

func TestGetSessionData(t *testing.T) {
	suite := NewTestSuite(t)
	expectedSessionData := map[string]any{"test_key": "test_value"}
	ctx := context.Background()
	testCases := []testCase{
		{
			name: "success",
			mock: func() {
				storedSessionData, err := json.Marshal(expectedSessionData)
				require.NoError(t, err)
				suite.mockStorage.EXPECT().Get(ctx, suite.sessionManager.StorageKey()).Return(string(storedSessionData), nil)
			},
			expected: expectedSessionData,
			err:      nil,
		},
		{
			name: "storage get error",
			mock: func() {
				suite.mockStorage.EXPECT().Get(ctx, suite.sessionManager.StorageKey()).Return("", fakeErr)
			},
			expected: map[string]any(nil),
			err:      fakeErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			data, err := suite.sessionManager.GetSessionData(ctx)
			assert.Equal(t, tc.expected, data)
			assert.ErrorIs(t, err, tc.err)
		})
	}
}

func TestSetSessionData(t *testing.T) {
	suite := NewTestSuite(t)
	ctx := context.Background()
	existentSessionData := map[string]any{"test_key": "test_value"}
	expectedKey := gofakeit.Name()
	expectedVal := gofakeit.UUID()
	testCases := []testCase{
		{
			name: "success",
			mock: func() {
				rawData, err := json.Marshal(existentSessionData)
				require.NoError(t, err)
				suite.mockStorage.EXPECT().Get(ctx, suite.sessionManager.StorageKey()).Return(string(rawData), nil)
				updatedSessionData := map[string]any{"test_key": "test_value", expectedKey: expectedVal}
				marshalledUpdatedData, err := json.Marshal(updatedSessionData)
				require.NoError(t, err)
				suite.mockStorage.EXPECT().Set(ctx, suite.sessionManager.StorageKey(), string(marshalledUpdatedData)).Return(nil)
			},
		},
		{
			name: "storage get error",
			mock: func() {
				suite.mockStorage.EXPECT().Get(ctx, suite.sessionManager.StorageKey()).Return("", fakeErr)
			},
			err: fakeErr,
		},
		{
			name: "storage set error",
			mock: func() {
				rawData, err := json.Marshal(existentSessionData)
				require.NoError(t, err)
				suite.mockStorage.EXPECT().Get(ctx, suite.sessionManager.StorageKey()).Return(string(rawData), nil)
				updatedSessionData := map[string]any{"test_key": "test_value", expectedKey: expectedVal}
				marshalledUpdatedData, err := json.Marshal(updatedSessionData)
				require.NoError(t, err)
				suite.mockStorage.EXPECT().Set(ctx, suite.sessionManager.StorageKey(), string(marshalledUpdatedData)).Return(fakeErr)
			},
			err: fakeErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := suite.sessionManager.SetToSession(context.Background(), expectedKey, expectedVal)
			assert.ErrorIs(t, err, tc.err)
		})
	}

}
