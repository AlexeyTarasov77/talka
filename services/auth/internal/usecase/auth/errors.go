package auth

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrOAuthSignInFailed     = errors.New("oauth signin failed")
	ErrOAuthSignInNotStarted = errors.New("oauth signin process was not started to complete")
)
