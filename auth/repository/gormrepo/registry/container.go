package registry

import (
	"github.com/dannywolfmx/ReSender/auth/domain/model"
	"github.com/dannywolfmx/ReSender/auth/domain/service"
	"github.com/dannywolfmx/ReSender/auth/repository/gormrepo"
	"github.com/dannywolfmx/ReSender/auth/usecase"
	"github.com/jinzhu/gorm"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	dataBaseFile := "/home/daniel/Programacion/ReSender/db/data/data.db"

	build, err := di.NewBuilder()

	if err != nil {
		return nil, err
	}

	err = build.Add([]di.Def{
		{
			Name: "gormSqlite",
			Build: func(ctn di.Container) (interface{}, error) {
				db, err := gorm.Open("sqlite3", dataBaseFile)

				if err != nil {
					return nil, err
				}

				db.AutoMigrate(&model.User{})

				return db, err

			},
			Close: func(db interface{}) error {
				return db.(*gorm.DB).Close()
			},
		},
		{
			Name: "usecase",
			Build: func(ctn di.Container) (interface{}, error) {
				connDB := ctn.Get("gormSqlite").(*gorm.DB)
				repo := gormrepo.NewUserRepository(connDB)
				service := service.NewUserService(repo)
				return usecase.NewAuthUsecase(repo, service), nil
			},
		},
	}...)

	if err != nil {
		return nil, err
	}

	return &Container{
		ctn: build.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}
