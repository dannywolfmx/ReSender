package auth

import "errors"

var ErrInvalidToken = errors.New("invalid access token")
var ErrInvalidPassword = errors.New("invalid password")
var ErrNameAlreayExist = errors.New("the username already exist")
