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
	Number   string          `json:"number"`
	Invoice  string          `json:"invoice"`
	ClientID uint            `json:"client_id"`
	Mails    []MailDirection `json:"mails" gorm:"many2many:senders"`
	Files    []File          `json:"files"`
}

type Client struct {
	Orm    `json:"orm"`
	Name   string `json:"name"`
	Orders []Order
}

//TODO: Estructuras
type MailDirection struct {
	Orm       `json:"orm"`
	Direction string `json:"direction"`
}

type File struct {
	Orm     `json:"orm"`
	Path    string `json:"path"`
	Title   string `json:"title"`
	OrderID uint   `json:"order_id"`
}
