package config

import "errors"

var (
	ErrUserNotFound    = errors.New("user was not found")
	ErrUserInsertError = errors.New("user could not be inserted")
	ErrUserExists      = errors.New("user already exists")
	ErrInvalidUser     = errors.New("user is invalid")
	ErrInvalidEmail    = errors.New("email is invalid")
	ErrInvalidPass     = errors.New("password is invalid")
)
