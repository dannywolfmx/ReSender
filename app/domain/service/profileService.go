package service

import (
	"github.com/dannywolfmx/ReSender/app"
	"golang.org/x/crypto/bcrypt"
)

//Coloca aqui toda la logca de negocio, recuerda que este tipo de logica no debe tener persistencia.

//Nota ClientService debe ser publico, dado que el usecase hace uso de este tipo
type ProfileService struct {
	repo app.ProfileRepository
}

//NewProfileService is a contructor to get a good formed ProfileService
func NewProfileService(repo app.ProfileRepository) *ProfileService {
	return &ProfileService{
		repo: repo,
	}
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

//ComparePasswordHash check if the password is the same like the hash
//Example taked from: https://gowebexamples.com/password-hashing/
func (p *ProfileService) ComparePasswordHash(pass, hash string) bool {
	// TODO Change the bcrypt.MinCost constant
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	//If err is nil, the password works with the hash
	//if err is another value, the password is wrong
	return err == nil

}
