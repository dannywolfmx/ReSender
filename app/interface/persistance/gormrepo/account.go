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

func (a *accountRepository) Save(m *model.Account) error {
	a.db.Save(m)
	return nil
}

func (a *accountRepository) GetById(id uint) model.Account {
	account := model.Account{}
	a.db.Where("id = ?", id).First(account)
	return account
}

func (a *accountRepository) All() ([]model.Account, error) {
	accounts := []model.Account{}
	a.db.Preload("MailConfig").Find(&accounts)
	return accounts, nil
}

func (a *accountRepository) Detele(id uint) error {
	a.db.Where("id = ? ", id).Delete(&model.Account{})
	return nil
}

func (a *accountRepository) Update(u *model.Account) error {
	a.db.Model(u).Update(u)
	return nil
}
