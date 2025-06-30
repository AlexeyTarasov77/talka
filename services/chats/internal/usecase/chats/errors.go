package chats

import "errors"

var (
	ErrChatAlreadyExists   = errors.New("chat already exists")
	ErrChatNotFound        = errors.New("chat is not found")
	ErrUserNotFound        = errors.New("related user not found")
	ErrMemberNotFound      = errors.New("One or more of provided members was not found")
	ErrMemberAlreadyInChat = errors.New("member already in chat")
)
