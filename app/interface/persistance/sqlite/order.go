package sqlite

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)

type orderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Save(m *model.Order) error {
	query := ` INSERT INTO orden (
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
	order := new(model.Order)
	query := `
		SELECT * from orden LIMIT 1 WHERE invoice = $1
	`
	err := r.db.Get(order, query, invoice)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) All() ([]*model.Order, error) {
	query := `SELECT * FROM orden`
	orders := []*model.Order{}
	err := r.db.Select(&orders, query)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
