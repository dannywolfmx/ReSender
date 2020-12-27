package registry

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sarulabs/di"

	appModel "github.com/dannywolfmx/ReSender/app/domain/model"
	appService "github.com/dannywolfmx/ReSender/app/domain/service"
	appRepository "github.com/dannywolfmx/ReSender/app/repository/gormrepo"
	appUsecase "github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/dannywolfmx/ReSender/auth"

	authModel "github.com/dannywolfmx/ReSender/auth/domain/model"
	authService "github.com/dannywolfmx/ReSender/auth/domain/service"
	authRepository "github.com/dannywolfmx/ReSender/auth/repository/gormrepo"
	authUsecase "github.com/dannywolfmx/ReSender/auth/usecase"
)

type Container struct {
	ctn di.Container
}

type DIContainer struct{
	AuthUsecase auth.AuthUsecase
}

func NewDIContainer(dbType string) (*DIContainer, error){

	connDB, err:= gormSqliteConnection("db/data/data.db")
	if err != nil{
		return nil, err
	}

	authUsecase, err := authUsecaseGORM(connDB)
	if err != nil{
		return nil, err
	}

	diContainer := &DIContainer{
		AuthUsecase: authUsecase,
	}

	return diContainer, nil
}


//TODO delete dependency to *gorm.DB type, instance use own interface
//Close the connection when the app close
func authUsecaseGORM(connDB *gorm.DB) (auth.AuthUsecase, error){
	repo := authRepository.NewUserRepository(connDB)
	service := authService.NewUserService(repo)
	return authUsecase.NewAuthUsecase(repo, service), nil
}

func gormSqliteConnection(dataBasePath string) (*gorm.DB, error){
		//Fijar ruta de la db y el tipo de db
		db, err := gorm.Open("sqlite3", dataBasePath)
		if err != nil {
			return nil, err
		}
		//Migrar base de datos para hacer match con estructuras
		//TODO Leer desde una variable de entorno si es que debemos migrar la db
		//con el motivo de que solo hacerlo cuando sea necesario
		migrarDBGorm(db)
		return db, nil
}


func NewContainer() (*Container, error) {
	dataBaseFile := "/home/daniel/Programacion/ReSender/db/data/data.db"

	//dataBaseFile := "/tmp/data.db"
	build, err := di.NewBuilder()

	if err != nil {
		return nil, err
	}

	err = build.Add([]di.Def{
		{
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
				repo := appRepository.NewOrderRepository(connDB)
				service := appService.NewOrderService(repo)
				return appUsecase.NewOrderUsecase(repo, service), nil
			},
		}, {
			Name: "client-usecase",
			Build: func(ctn di.Container) (interface{}, error) {
				connDB := ctn.Get("gormSqlite").(*gorm.DB)
				//TODO: Revisar por que no devuelvo un puntero en el repositorio
				repo := appRepository.NewClientRepository(connDB)
				service := appService.NewClientService(repo)
				return appUsecase.NewClientUsecase(repo, service), nil
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
				repository := appRepository.NewProfileRepository(connDB)

				//Servicio del profile
				//El servicio requiere hacer operaciones con el repositorio por lo que se envia uno funcional
				//Dependiedo de la funcionalidad puede compartir el mismo repo que el usecase
				service := appService.NewProfileService(repository)

				//Cramos un usecase con un repositorio y un repositorio
				return appUsecase.NewProfileUsecase(repository, service), nil
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
		//App models tables
		&appModel.Order{},
		&appModel.Client{},
		&appModel.MailDirection{},
		&appModel.File{},
		&appModel.Profile{},
		&appModel.MailServer{},

		//Auth models tables
		&authModel.User{},
	)
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}
