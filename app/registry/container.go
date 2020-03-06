package registry

import (
	"github.com/dannywolfmx/ReSender/app/domain/service"
	"github.com/dannywolfmx/ReSender/app/interface/persistance"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/dannywolfmx/ReSender/db"
	"github.com/jmoiron/sqlx"
	"github.com/sarulabs/di"
)

type container struct {
	ctn di.Container
}

func NewContainer() (*container, error) {
	build, err := di.NewBuilder()

	if err != nil {
		return nil, err
	}

	err = build.Add([]di.Def{{
		Name: "sqlite",
		Build: func(ctn di.Container) (interface{}, error) {
			pathDB := "./db/data/data.db"
			db.DB, err = db.NewDBSqliteConnection(pathDB).InitDB()
			return db.DB, err
		},
	}, {
		Name: "order-usercase",
		Build: func(ctn di.Container) (interface{}, error) {
			connDB := ctn.Get("sqlite").(*sqlx.DB)
			repo := persistance.NewOrderRepository(connDB)
			service := service.NewOrderService(repo)
			return usecase.NewOrderUsecase(repo, service), nil
		},
	}}...)

	if err != nil {
		return nil, err
	}

	return &container{
		ctn: build.Build(),
	}, nil
}
