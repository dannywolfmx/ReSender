package registry

import (
	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/domain/service"
	"github.com/dannywolfmx/ReSender/app/interface/persistance/gormrepo"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	dataBaseFile := "/home/daniel/Programacion/ReSender/db/data/data.db"

	//dataBaseFile := "/tmp/data.db"
	build, err := di.NewBuilder()

	if err != nil {
		return nil, err
	}

	err = build.Add([]di.Def{{
		Name: "gormSqlite",
		Build: func(ctn di.Container) (interface{}, error) {
			//Fijar ruta de la db y el tipo de db
			db, err := gorm.Open("sqlite3", dataBaseFile)
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
		Name: "order-usecase",
		Build: func(ctn di.Container) (interface{}, error) {
			//connDB := ctn.Get("sqlite").(*sqlx.DB)
			connDB := ctn.Get("gormSqlite").(*gorm.DB)
			repo := gormrepo.NewOrderRepository(connDB)
			service := service.NewOrderService(repo)
			return usecase.NewOrderUsecase(repo, service), nil
		},
	}, {
		Name: "client-usecase",
		Build: func(ctn di.Container) (interface{}, error) {
			connDB := ctn.Get("gormSqlite").(*gorm.DB)
			//TODO: Revisar por que no devuelvo un puntero en el repositorio
			repo := gormrepo.NewClientRepository(connDB)
			service := service.NewClientService(repo)
			return usecase.NewClientUsecase(repo, service), nil
		},
	}, {
		// CONTENEDOR PARA PROFILE
		//Nombre del contenedor
		Name: "profile-usecase",
		//Forma en que se va a construir el contenedor
		Build: func(ctn di.Container) (interface{}, error) {

			//Base de datos seleccionada
			connDB := ctn.Get("gormSqlite").(*gorm.DB)

			//TODO: Revisar por que no devuelvo un puntero en el repositorio
			//Repositorio del profile
			repository := gormrepo.NewProfileRepository(connDB)

			//Servicio del profile
			//El servicio requiere hacer operaciones con el repositorio por lo que se envia uno funcional
			//Dependiedo de la funcionalidad puede compartir el mismo repo que el usecase
			service := service.NewProfileService(repository)

			//Cramos un usecase con un repositorio y un repositorio
			return usecase.NewProfileUsecase(repository, service), nil
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

//migrarDBGorm se encarga de realizar el proceso de migracion de las tablas del modelo
func migrarDBGorm(db *gorm.DB) {
	//Migrar estructuras
	db.AutoMigrate(
		&model.Order{},
		&model.Client{},
		&model.MailDirection{},
		&model.File{},
		&model.Profile{},
		&model.MailServer{},
	)
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}
