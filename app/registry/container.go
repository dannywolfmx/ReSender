package registry

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/service"
	"github.com/dannywolfmx/ReSender/app/interface/persistance/gormrepo"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/dannywolfmx/ReSender/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	build, err := di.NewBuilder()

	if err != nil {
		return nil, err
	}

	err = build.Add([]di.Def{{
		Name: "sqlite",
		Build: func(ctn di.Container) (interface{}, error) {
			pathDB := "./db/data/data.db"
			//TODO: Cambiar db.DB a una variable local
			return db.NewDBSqliteConnection(pathDB).InitDB()
		},
	}, {
		Name: "gormSqlite",
		Build: func(ctn di.Container) (interface{}, error) {
			//Fijar ruta de la db y el tipo de db
			db, err := gorm.Open("sqlite3", "./db/data/data.db")
			if err != nil {
				return nil, err
			}
			//Migrar base de datos para hacer match con estructuras
			migrarDBGorm(db)
			return db, err
		},
		Close: func(db interface{}) error {
			return db.(*gorm.DB).Close()
		},
	}, {
		Name: "order-usercase",
		Build: func(ctn di.Container) (interface{}, error) {
			//connDB := ctn.Get("sqlite").(*sqlx.DB)
			connDB := ctn.Get("gormSqlite").(*gorm.DB)
			repo := gormrepo.NewOrderRepository(connDB)
			service := service.NewOrderService(repo)
			return usecase.NewOrderUsecase(repo, service), nil
		},
	}}...)

	if err != nil {
		return nil, err
	}

	return &Container{
		ctn: build.Build(),
	}, nil
}

//migrarDBGorm se encarga de realizar el proceso de migracion de las tablas del modelo
func migrarDBGorm(db *gorm.DB) {
	//Migrar estructuras
	db.AutoMigrate(&model.Order{}, &model.Client{})
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}
