package entity

import "errors"

var (
	ErrPasswordIsNotValid    = errors.New("password must have from 8 to 30 characters")
	ErrPasswordIsNotStrong   = errors.New("password is not strong")
	ErrEmailIsNotValid       = errors.New("email is not valid")
	ErrEmailHasExisted       = errors.New("email has existed")
	ErrLoginFailed           = errors.New("email and password are not valid")
	ErrFirstNameIsEmpty      = errors.New("first name can not be blank")
	ErrRefreshTokenIsEmpty   = errors.New("refresh token can not empty")
	ErrFirstNameTooLong      = errors.New("first name too long, max character is 30")
	ErrLastNameIsEmpty       = errors.New("last name can not be blank")
	ErrLastNameTooLong       = errors.New("last name too long, max character is 30")
	ErrCannotRegister        = errors.New("cannot register")
	ErrorValidateFailed      = errors.New("validation failed")
	ErrorRefreshTokenFailed  = errors.New("refresh token failed")
	ErrorRefreshTokenExpired = errors.New("refresh token was expired")
)
