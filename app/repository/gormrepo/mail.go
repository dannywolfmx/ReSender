package gormrepo

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/jinzhu/gorm"
)

//mailRepository repositorio hecho en gorm.
//para ser usado se debe usar NewMailRepository
type mailRepository struct {
	db *gorm.DB
}

func NewMailRepository(db *gorm.DB) *mailRepository {
	return &mailRepository{
		db: db,
	}
}

//Find a mail direction, ej. "test@linux.com"
func (m *mailRepository) Find(direction string) (*model.MailDirection, error) {
	mail := new(model.MailDirection)
	m.db.Where("direction = ?", direction).First(mail)
	return mail, nil
}

//Delete a mail with the id
func (m *mailRepository) Detele(id uint) error {
	m.db.Where("id = ? ", id).Delete(&model.MailDirection{})
	return nil
}

//DeleteByAddress delete a mail using his mail address
func (m *mailRepository) DeleteByAddress(direction string) error {
	m.db.Where("direction = ? ", direction).Delete(&model.MailDirection{})
	return nil
}

//Update a mail direction
func (m *mailRepository) Update(mail *model.MailDirection) error {
	m.db.Model(mail).Update(mail)
	return nil
}
