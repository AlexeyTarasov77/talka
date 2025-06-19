package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/AlexeyTarasov77/messanger.users/internal/dto"
	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/storage"
	"github.com/AlexeyTarasov77/messanger.users/internal/usecase/auth"
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
