package persistence

import (
	"sync"

	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)

type orderRepository struct {
	mu *sync.Mutex
	db *sqlx.DB
}

func (r *orderRepository) Save(m *model.Order) error {
	query := `
		INSERT INTO order (
			id,
			number,
			invoice
		) Values ($1,$2,$3)
	`
	//Insertar valores en la tabla
	_, err := r.db.Exec(query, xid.New(), m.Number, m.Invoice)
	return err
}

func (r *orderRepository) FindByInvoice(invoice string) (*model.Order, error) {
	return nil, nil
}

func (r *orderRepository) All() ([]*model.Order, error) {
	query := `SELECT * FROM order`
	orders := []*model.Order{}
	r.db.Select(orders, query)
	return orders, nil
}
