package config

import "errors"

var (
	ErrUserNotFound         = errors.New("user was not found")
	ErrUserInsertError      = errors.New("user could not be inserted")
	ErrUserExists           = errors.New("user already exists")
	ErrInvalidUser          = errors.New("user is invalid")
	ErrInvalidEmail         = errors.New("email is invalid")
	ErrInvalidPassword      = errors.New("password is invalid")
	ErrURLNotFound          = errors.New("url was not found")
	ErrURLsNotFound         = errors.New("urls were not found")
	ErrSessionCreation      = errors.New("session could not be created")
	ErrSessionNotFound      = errors.New("session was not found")
	ErrFailedToReadDocument = errors.New("failed to read document")
)
