package service

import (
	"fmt"

	"github.com/dannywolfmx/ReSender/auth"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo auth.UserRepository
}

func NewUserService(repo auth.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

//Duplicated check if the name is already registred
func (p *UserService) Duplicated(name string) error {
	//If the profile dont exist you will get a nil pointer
	user, err := p.repo.Get(name)
	if err != nil {
		return err
	}

	if user != nil {
		return fmt.Errorf("El nombre ya esta registrado %s", user.Username)
	}

	return nil
}

//HashAndSaltPassword get a hashed and salted password
//Example taked from: https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func (p *UserService) HashAndSaltPassword(pass string) (string, error) {
	// TODO Change the bcrypt.MinCost constant
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//ComparePasswordHash check if the password is the same like the hash
//Example taked from: https://gowebexamples.com/password-hashing/
func (p *UserService) ComparePasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	//If err is nil, the password works with the hash
	//if err is another value, the password is wrong
	return err == nil
}
