package internalErrors

import "errors"

var (
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidPwd           = errors.New("invalid password")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrUserNotFound         = errors.New("user not found")
	ErrTokenExpired         = errors.New("token expired")
	ErrTokenInvalid         = errors.New("token invalid")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrCacheItemNotFound    = errors.New("cache item not found")
	ErrCacheItemExpired     = errors.New("cache item expired")
)
