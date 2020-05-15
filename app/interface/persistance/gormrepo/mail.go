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

func (a *accountRepository) Save(_ *model.Account) error {
	panic("not implemented") // TODO: Implement
}

func (a *accountRepository) GetById(id uint) model.Account {
	panic("not implemented") // TODO: Implement
}

func (a *accountRepository) All() ([]model.Account, error) {
	panic("not implemented") // TODO: Implement
}

func (a *accountRepository) Detele(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (a *accountRepository) Update(_ *model.Account) error {
	panic("not implemented") // TODO: Implement
}
