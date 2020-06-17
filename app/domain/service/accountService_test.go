package service

import "testing"

func TestHashAndSaltPassword(t *testing.T) {
	password := []byte("123abc")

	hashedPassword, err := HashAndSaltPassword(password)

	//El hash no se completo de buena forma
	if err != nil {
		t.Error("error al realizar el Hash en la passowrd", err)
	}

	t.Log("Resultado exitoso del hash: ", hashedPassword)
}
