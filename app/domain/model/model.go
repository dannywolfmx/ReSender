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
	Orm
	Number  string
	Invoice string
	Clients *Client

	//	Mails   []MailDirection
	//	Files   []File
}

type Client struct {
	Orm
	Name string
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
