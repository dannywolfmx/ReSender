package service

import (
	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

//Coloca aqui toda la logca de negocio, recuerda que este tipo de logica no debe tener persistencia, asi que las llamadas a una base de datos, no van aqui.

//Nota ClientService debe ser publico, dado que el usecase hace uso de este tipo
type ProfileService struct {
	repo repository.Profile
}

//NewProfileService is a contructor to get a good formed ProfileService
func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{
		repo: repo,
	}
}

//HashAndSaltPassword get a hashed and salted password
//Example taked from: https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func HashAndSaltPassword(pass []byte) (string, error) {
	// TODO Change the bcrypt.MinCost constant
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
