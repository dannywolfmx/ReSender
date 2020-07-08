package service

import (
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
