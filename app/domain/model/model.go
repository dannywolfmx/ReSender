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
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

//Order estructura que almacena la informacion de una orden de compra
type Order struct {
	Orm      `json:"orm"`
	ClientID uint `json:"client_id" validate:"required"`

	//Estado representa el estatus en el que se encuentra esta orden
	// 0 - No enviada {ESTADO DEFAULT}
	// 1 - Enviada
	// TODO - Algunos clientes no tienen la necesidad de este estado, pendiente de encontrar una mejor solución
	// 2 - Pendiente de orden de entrada a almacen
	Estado uint `json:"estado" validate:"required"`

	//Files archivos relacionados con esta orden de compra
	Files []*File `json:"files"`

	//Invoice representa al numero de la factura
	Invoice string `json:"invoice"`

	//Mails lista de direcciones de correo
	Mails []*MailDirection `json:"mails" gorm:"many2many:senders"`

	//Number representa al numero de la orden de compra
	Number string `json:"number"`
}

//Client es una estructura que almacena la informacion personal de este cliente asi como su relacion con las ordenes
type Client struct {
	Orm `json:"orm"`

	//ProfiletID Id del perfil al que pertenece este cliente
	ProfiletID uint `json:"profilet_id" validate:"required"`

	//Name nombre del cliente
	Name string `json:"name" validate:"required"`

	//Order ordenes relacionadas con este cliente
	Orders []*Order `json:"orders"`
}

//MailDirection es una estructura que almacena la informacion que un correo electrónico pueda tener
type MailDirection struct {
	Orm `json:"orm"`

	//Direction direccion de correo electronico. Ej "prueba@gmail.com"
	Direction string `json:"direction" validate:"required,email"`
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

//Profile es una estructura que almacena la informacion de un usuario
type Profile struct {
	Orm

	//Clients lista de clientes que estan asociados con el perfil
	Clients []*Client

	//DefailtMailConfigId la configuracion que esta marcada por defecto
	//Nota: esto no esta relacionado con GORM
	DefaultMailConfigID uint

	//ImageAvatarPath Imagen del perfil
	ImageAvatarPath string

	//MailConfig configuraciones asociadas al perfil del usuario
	MailConfig []*MailServer

	//Name nombre del perfil
	Name string

	//Password profile
	Password string
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
