package gormrepo

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/jinzhu/gorm"
)

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *fileRepository {
	return &fileRepository{
		db: db,
	}
}

func (f *fileRepository) Get(id uint) (*model.File, error) {
	file := new(model.File)
	f.db.Where("id = ?", id).First(file)
	return file, nil
}

func (f *fileRepository) Detele(id uint) error {
	f.db.Where("id = ?", id).Delete(&model.File{})
	return nil
}

func (f *fileRepository) Update(file *model.File) error {
	f.db.Model(file).Update(file)
	return nil
}
