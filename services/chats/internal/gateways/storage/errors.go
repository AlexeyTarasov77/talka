package storage

import "errors"

var (
	ErrNotFound      = errors.New("record not found")
	ErrAlreadyExists = errors.New("record already exists")
	// ErrRelationNotFound indicates foreign key violation
	ErrRelationNotFound = errors.New("related record not found")
)
