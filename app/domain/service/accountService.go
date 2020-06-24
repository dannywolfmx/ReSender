package service

import (
	"fmt"
	"log"

	"github.com/dannywolfmx/ReSender/app/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

//Coloca aqui toda la logca de negocio, recuerda que este tipo de logica no debe tener persistencia.

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

//Duplicated check if the name is already registred
func (p *ProfileService) Duplicated(name string) error {
	//If the profile dont exist you will get a nil pointer
	profile, err := p.repo.GetByName(name)
	if err != nil {
		return err
	}

	log.Println(name)

	if profile != nil {
		return fmt.Errorf("El nombre ya esta registrado %s", profile.Name)
	}

	return nil
}

//HashAndSaltPassword get a hashed and salted password
//Example taked from: https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func (p *ProfileService) HashAndSaltPassword(pass string) (string, error) {
	// TODO Change the bcrypt.MinCost constant
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
