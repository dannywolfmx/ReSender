package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string
	Password string
}

type UserClaims struct {
	jwt.StandardClaims
	User *User
}
