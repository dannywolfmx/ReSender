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

func (r *orderRepository) Save(m *model.Order) error {
	r.db.Create(m)
	return nil
}

func (r *orderRepository) FindByInvoice(invoice string) (*model.Order, error) {
	order := new(model.Order)
	r.db.Where(&model.Order{Invoice: invoice}).First(order)
	return order, nil
}

func (r *orderRepository) All() ([]model.Order, error) {
	orders := []model.Order{}
	r.db.Find(&orders)

	return orders, nil
}

//Find and Delete all the matches record
//Note: Delete is a soft delete, this function just set a flag
//You need to use r.db.Unscoped().Delete(&model.Order{}) to clear the Delete records permanently
func (r *orderRepository) Detele(invoice string) error {
	r.db.Where("invoice = ?", invoice).Delete(&model.Order{})
	return nil
}

func (r *orderRepository) Update(order *model.Order) error {
	//Save will update all the fields, even it is not changed
	r.db.Save(order)
	return nil
}
