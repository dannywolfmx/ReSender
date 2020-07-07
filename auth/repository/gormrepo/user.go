package gormrepo

import (
	"github.com/dannywolfmx/ReSender/auth/domain/model"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *model.User) error {
	r.db.Save(user)
	return nil
}

func (r *userRepository) Get(username string, password string) *model.User {
	user := &model.User{}
	r.db.Where("username = ?", username).Find(user)
	return user
}
