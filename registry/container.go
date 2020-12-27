package registry

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/dannywolfmx/ReSender/app"
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

type DIContainer struct{
	AuthUsecase auth.AuthUsecase
	OrderUsecase app.OrderUsecase
	ProfileUsecase app.ProfileUsecase
	ClientUsecase app.ClientUsecase
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
	orderUsecase, err := orderUsecaseGORM(connDB)
	if err != nil{
		return nil, err
	}
	profileUsecase, err := profileUsecaseGORM(connDB)
	if err != nil{
		return nil, err
	}
	clientUsecase, err := clientUsecaseGORM(connDB)
	if err != nil{
		return nil, err
	}

	diContainer := &DIContainer{
		AuthUsecase: authUsecase,
		OrderUsecase: orderUsecase,
		ProfileUsecase: profileUsecase,
		ClientUsecase: clientUsecase,
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

func orderUsecaseGORM(connDB *gorm.DB) (app.OrderUsecase, error){
	repo := appRepository.NewOrderRepository(connDB)
	service := appService.NewOrderService(repo)
	return appUsecase.NewOrderUsecase(repo, service), nil
}

func clientUsecaseGORM(connDB *gorm.DB) (app.ClientUsecase, error){
	repo := appRepository.NewClientRepository(connDB)
	service := appService.NewClientService(repo)
	return appUsecase.NewClientUsecase(repo, service), nil
}

func profileUsecaseGORM(connDB *gorm.DB) (app.ProfileUsecase, error){
	repository := appRepository.NewProfileRepository(connDB)
	service := appService.NewProfileService(repository)
	return appUsecase.NewProfileUsecase(repository, service), nil
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
