package gormrepo

import (
	"github.com/dannywolfmx/ReSender/auth/domain/model"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{
		db: db,
	}
}

func (r userRepository) Create(user *model.User) error {
	r.db.Save(user)
	return nil
}

func (r userRepository) Get(username string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Where("username = ?", username).Find(user).Error
	if gorm.IsRecordNotFoundError(err) {
		//No record and no error
		return nil, nil
	}
	return user, err
}
