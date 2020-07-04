package gormrepo

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/jinzhu/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

//Find and Delete all the matches record
//Note: Delete is a soft delete, this function just set a flag
//You need to use r.db.Unscoped().Delete(&model.Order{}) to clear the Delete records permanently
func (r *orderRepository) Detele(id uint) error {
	r.db.Where("id = ?", id).Delete(&model.Order{})
	return nil
}

func (r *orderRepository) Update(order *model.Order) error {
	//Save will update all the fields, even it is not changed
	r.db.Model(order).Update(order)
	return nil
}
