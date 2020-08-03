package model

import (
	"time"
)

//Orm es una estrucura para ser embebida no esta planeada para instanciarla sola

/***
Representación ideal de una estructura ideal en JSON

Profiles:[
   {
    DefailtMailConfigId:1,
    MailConfig:[
     Adress: "????",
     Alias: "correo de google",
     From:"prueba@gmail.com",
     Password:"123abc",
     Sever: "?????"
    ],
    Name:"Cuenta principal",
    ImageAvatarPath:"/home/daniel/imagen.png",
    Clients: [
       Client{
          Orders: [
              Order{
                Estado: 0,
                Number: "ABS123",
                Invoice: "Fac123",
                 Mails{
                    Direction: "prueba@gmail.com"
                 },
                 Files{
                    Path: "/home/daniel/archivo/1.pdf",
                    Title: "Archivo de prueba",
                 }
           }
         ]
       }
    ]
   }
]

***/
//Orm almacena los metadatos de una estructura.
//Es una estructura que emula der la de GORM con el motivo de no depender de la api del framework
type Orm struct {
	ID uint `gorm:"primary_key" json:"id"`
	//No necesitamos enviar esto en formato json, por lo que se omiten
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//Order estructura que almacena la informacion de una orden de compra
type Order struct {
	Orm
	ClientID uint

	//Estado representa el estatus en el que se encuentra esta orden
	// 0 - No enviada {ESTADO DEFAULT}
	// 1 - Enviada
	// TODO - Algunos clientes no tienen la necesidad de este estado, pendiente de encontrar una mejor solución
	// 2 - Pendiente de orden de entrada a almacen
	Estado *uint

	//Files archivos relacionados con esta orden de compra
	Files []*File

	//Invoice representa al numero de la factura
	Invoice string

	//Mails lista de direcciones de correo
	Mails []*MailDirection

	//Number representa al numero de la orden de compra
	Number string
}

func NewOrder(clientID uint, estado *uint, invoice, number string) *Order {
	return &Order{
		ClientID: clientID,
		Estado:   estado,
		Invoice:  invoice,
		Number:   number,
	}
}

//Client es una estructura que almacena la informacion personal de este cliente asi como su relacion con las ordenes
type Client struct {
	Orm

	//ProfiletID Id del perfil al que pertenece este cliente
	ProfiletID uint

	//Name nombre del cliente
	Name string

	//Order ordenes relacionadas con este cliente
	Orders []*Order
}

func NewClient(profileID uint, name string) *Client {
	return &Client{
		ProfiletID: profileID,
		Name:       name,
	}
}

//MailDirection es una estructura que almacena la informacion que un correo electrónico pueda tener
type MailDirection struct {
	Orm

	//Direction direccion de correo electronico. Ej "prueba@gmail.com"
	Direction string
}

func NewMailDirection(direction string) *MailDirection {
	return &MailDirection{
		Direction: direction,
	}
}

//File es una estructura que almacena la metadata y localización de un archivo
type File struct {
	Orm
	OrderID uint

	//Path direccion en la que se encuentra almacenado el archivo
	Path string

	//Title titulo del archivo
	Title string
}

func NewFile(path, title string) *File {
	return &File{
		Path:  path,
		Title: title,
	}
}

//Profile es una estructura que almacena la informacion de un usuario
type Profile struct {
	Orm

	//Clients lista de clientes que estan asociados con el perfil
	Clients []*Client

	//DefailtMailConfigId la configuracion que esta marcada por defecto
	//Nota: esto no esta relacionado con GORM
	//DefailtMailConfigID puede ser nil, por ello se coloca como puntero
	DefaultMailConfigID *uint

	//ImageAvatarPath Imagen del perfil
	ImageAvatarPath string

	//MailConfig configuraciones asociadas al perfil del usuario
	MailConfig []*MailServer

	UserID uint
}

//MailServer almacena la informacion de configuracion de un servidor de correos
type MailServer struct {
	Orm

	//Relationship
	ProfiletID uint

	//Address represent the server direction including the port.
	//Ej "mail.example.com:25"
	Address string

	//Alias de este perfil del servidor
	Alias string

	//From describe el correo electronico del remitente
	From string

	//Password password del servidor de mail
	//TODO  verificar si se realiza algun tipo de encriptacion
	Password string

	//Server ??? //Creo que es server name
	Server string
}

//NewMailServer create a new MailServer
func NewMailServer(profileID uint, address, alias, from, password, server string) *MailServer {
	return &MailServer{
		ProfiletID: profileID,
		Address:    address,
		Alias:      alias,
		From:       from,
		Password:   password,
		Server:     server,
	}
}
