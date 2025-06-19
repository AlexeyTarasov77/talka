package repo_test

import (
	"context"
	"testing"

	postgres_repos "github.com/AlexeyTarasov77/messanger.users/internal/gateways/storage/postgres"
	"github.com/AlexeyTarasov77/messanger.users/internal/usecase"
	"github.com/AlexeyTarasov77/messanger.users/pkg/postgres"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func newTestSuite(t *testing.T) (*postgres_repos.Repositories, context.Context) {
	ctx := context.Background()
	pg, tx := postgres.NewTestSuite(t, ctx)
	repos := postgres_repos.NewRepositorories(pg)
	ctx = usecase.SetTransaction(ctx, tx)
	return repos, ctx
}

func TestGetQueryableWithTx(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConnProvider := NewMockSupportsAcquire(ctrl)
	expectedTx := NewMockQueryable(ctrl)
	ctx := context.WithValue(context.Background(), usecase.TransactionCtxKey, expectedTx)
	queryable, err := postgres_repos.GetQueryable(ctx, mockConnProvider)
	assert.NoError(t, err)
	assert.Equal(t, expectedTx, queryable)
}

func TestGetQueryableNewConn(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	expectedConn := NewMockQueryable(ctrl)
	mockConnProvider := NewMockSupportsAcquire(ctrl)
	mockConnProvider.EXPECT().Acquire(context.Background()).Return(expectedConn, nil)
	queryable, err := postgres_repos.GetQueryable(ctx, mockConnProvider)
	assert.NoError(t, err)
	assert.Equal(t, expectedConn, queryable)
}
