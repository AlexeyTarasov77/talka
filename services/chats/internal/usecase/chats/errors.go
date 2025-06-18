package chats

import "errors"

var (
	ErrChatAlreadyExists = errors.New("chat already exists")
	ErrUserNotFound      = errors.New("related user not found")
	ErrMemberNotFound    = errors.New("One or more of provided members was not found")
)
