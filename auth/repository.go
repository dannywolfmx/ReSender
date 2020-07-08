package auth

import "github.com/dannywolfmx/ReSender/auth/domain/model"

//UserRepository
type UserRepository interface {
	Create(user *model.User) error
	Get(username string) (*model.User, error)
}
