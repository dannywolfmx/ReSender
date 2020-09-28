package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Orm struct {
	ID uint `gorm:"primary_key" json:"id"`
	//No necesitamos enviar esto en formato json, por lo que se omiten
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	Orm
	Username string
	Password string
}

type UserClaims struct {
	jwt.StandardClaims
	User *User
}
