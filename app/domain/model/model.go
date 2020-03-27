package model

import "time"

//Orm es una estrucura para ser embebida no esta planeada para instanciarla sola
type Orm struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Order struct {
	Orm      `json:"orm"`
	Number   string `json:"number"`
	Invoice  string `json:"invoice"`
	ClientID uint
	//	Mails   []MailDirection
	//	Files   []File
}

type Client struct {
	Orm    `json:"orm"`
	Name   string `json:"name"`
	Orders []Order
}

//TODO: Estructuras
type MailDirection struct {
	Orm
	Direction string
}

type File struct {
	Orm
	Path  string
	Title string
}
