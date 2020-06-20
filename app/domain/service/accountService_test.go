package service

import (
	"testing"

	"github.com/dannywolfmx/ReSender/app/domain/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestHashAndSaltPassword(t *testing.T) {
	password := []byte("123abc")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProfile(ctrl)
	serviceProfile := NewProfileService(m)

	hashedPassword, err := serviceProfile.HashAndSaltPassword(password)

	//El hash no se completo de buena forma
	if err != nil {
		t.Error("error al realizar el Hash en la passowrd", err)
	}

	t.Log("Resultado exitoso del hash: ", hashedPassword)
}

func TestDuplicated(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProfile(ctrl)

	m.EXPECT().GetByName("123abc").Return(nil, nil)

	serviceProfile := NewProfileService(m)

	err := serviceProfile.Duplicated("123abc")

	//El hash no se completo de buena forma

	t.Log("Resultado del duplicated: ", err)
}
