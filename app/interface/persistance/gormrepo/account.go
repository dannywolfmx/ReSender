package gormrepo

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/jinzhu/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{
		db: db,
	}
}

func (a *accountRepository) Save(m *model.Profile) error {
	a.db.Save(m)
	return nil
}

func (a *accountRepository) Get(id uint) (*model.Profile, error) {
	account := &model.Profile{}
	a.db.Where("id = ?", id).First(account)
	return account, nil
}

func (a *accountRepository) All() ([]*model.Profile, error) {
	accounts := []*model.Profile{}
	a.db.Preload("MailConfig").Find(&accounts)
	return accounts, nil
}

func (a *accountRepository) Detele(id uint) error {
	a.db.Where("id = ? ", id).Delete(&model.Profile{})
	return nil
}

func (a *accountRepository) Update(u *model.Profile) error {
	a.db.Model(u).Update(u)
	return nil
}
