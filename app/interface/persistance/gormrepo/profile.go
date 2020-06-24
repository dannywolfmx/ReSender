package gormrepo

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/jinzhu/gorm"
)

//profileRepository
type profileRepository struct {
	db *gorm.DB
}

//NewProfileRepository
func NewProfileRepository(db *gorm.DB) *profileRepository {
	return &profileRepository{
		db: db,
	}
}

//Save
func (p *profileRepository) Save(m *model.Profile) error {
	p.db.Save(m)
	return nil
}

//Get
func (p *profileRepository) Get(id uint) (*model.Profile, error) {
	profile := &model.Profile{}
	p.db.Where("id = ?", id).First(profile)
	return profile, nil
}

//All
func (p *profileRepository) All() ([]*model.Profile, error) {
	profile := []*model.Profile{}
	p.db.Preload("MailConfig").Find(&profile)
	return profile, nil
}

//Delete
func (p *profileRepository) Detele(id uint) error {
	p.db.Where("id = ? ", id).Delete(&model.Profile{})
	return nil
}

//Update
func (p *profileRepository) Update(u *model.Profile) error {
	p.db.Model(u).Update(u)
	return nil
}

//GetByName
func (p *profileRepository) GetByName(name string) (*model.Profile, error) {
	profile := &model.Profile{}
	err := p.db.Where("name = ?", name).First(profile).Error
	if gorm.IsRecordNotFoundError(err) {
		//No record and no error
		return nil, nil
	}
	return profile, err
}
