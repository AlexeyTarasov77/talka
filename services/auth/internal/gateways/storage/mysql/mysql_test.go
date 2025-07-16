package repo_test

import (
	"context"
	"testing"

	repo "github.com/AlexeyTarasov77/messanger.users/internal/gateways/storage/mysql"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/storage/mysql/sql/generated"
	"github.com/AlexeyTarasov77/messanger.users/internal/usecase"
	"github.com/AlexeyTarasov77/messanger.users/pkg/mysql"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// newDBTestSuite creates mysql test suite with opening new transaction then
// creates new context and binds opened transaction to it
func newDBTestSuite(t *testing.T) (*repo.Repositories, context.Context) {
	ctx := context.Background()
	suite := mysql.NewTestSuite(t).WithTx(ctx)
	repos := repo.NewRepositorories(suite.M)
	ctx = usecase.SetTransaction(ctx, suite.Tx)
	return repos, ctx
}

func TestGetQueriesWithTx(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTxFactory := NewMockTxFactory(ctrl)
	expectedTx := NewMockQueryableTransaction(ctrl)
	expectedQueries := generated.New(expectedTx)
	ctx := context.WithValue(context.Background(), usecase.TransactionCtxKey, expectedTx)
	baseRepo := repo.Base{mockTxFactory}

	queries, err := baseRepo.GetQueries(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedQueries, queries)
}

func TestGetQueryableNewConn(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockTxFactory := NewMockTxFactory(ctrl)
	baseRepo := repo.Base{mockTxFactory}
	expectedTx := NewMockQueryableTransaction(ctrl)
	expectedQueries := generated.New(expectedTx)
	mockTxFactory.EXPECT().Begin().Return(expectedTx, nil)

	queryies, err := baseRepo.GetQueries(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedQueries, queryies)
}
