package repository

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
)

type RepoOrder struct{}

func (r *RepoOrder) Save(order *model.Order) error {
	return nil
}

func (r *RepoOrder) All() ([]model.Order, error) {
	return nil, nil
}
