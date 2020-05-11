package gormrepo

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/jinzhu/gorm"
)

type mailRepository struct {
	db *gorm.DB
}

func NewMailRepository(db *gorm.DB) *mailRepository {
	return &mailRepository{
		db: db,
	}
}

func (r *mailRepository) Find(direction string) (*model.MailDirection, error) {
	mail := new(model.MailDirection)
	r.db.Where(&model.MailDirection{Direction: direction}).First(mail)
	return mail, nil
}

//Find and Delete all the matches record
//Note: Delete is a soft delete, this function just set a flag
//You need to use r.db.Unscoped().Delete(&model.Order{}) to clear the Delete records permanently
func (r *mailRepository) Detele(id uint) error {
	r.db.Where("id = ?", id).Delete(&model.MailDirection{})
	return nil
}

func (r *mailRepository) Update(mail *model.MailDirection) error {
	//Save will update all the fields, even it is not changed
	r.db.Model(mail).Update(mail)
	return nil
}
