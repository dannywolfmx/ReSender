package gormrepo

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/jinzhu/gorm"
)

//fileRepository
type fileRepository struct {
	db *gorm.DB
}

//NewFileRepository
func NewFileRepository(db *gorm.DB) *fileRepository {
	return &fileRepository{
		db: db,
	}
}

//Get
func (f *fileRepository) Get(id uint) (*model.File, error) {
	file := new(model.File)
	f.db.Where("id = ?", id).First(file)
	return file, nil
}

//Delete
func (f *fileRepository) Detele(id uint) error {
	f.db.Where("id = ?", id).Delete(&model.File{})
	return nil
}

//Update
func (f *fileRepository) Update(file *model.File) error {
	f.db.Model(file).Update(file)
	return nil
}
