package entity

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrGameNotFound = errors.New("game not found")
var ErrCommentNotFound = errors.New("comment not found")
var ErrCharacterNotFound = errors.New("character not found")
var ErrPeopleNotFound = errors.New("people not found")
var ErrInvalidEmail = errors.New("invalid Email")
var ErrInvalidPassword = errors.New("invalid Password")
var ErrEmptyAuthHeader = errors.New("empty auth header")
var ErrInvalidAuthHeader = errors.New("invalid auth header")
var ErrUserExists = errors.New("user with this email exists")
