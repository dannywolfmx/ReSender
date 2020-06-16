package model

import "time"

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
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//Order estructura que almacena la informacion de una orden de compra
type Order struct {
	Orm `json:"orm"`

	//Number representa al numero de la orden de compra
	Number string `json:"number"`

	//Invoice representa al numero de la factura
	Invoice  string `json:"invoice"`
	ClientID uint   `json:"client_id"`

	//Mails lista de direcciones de correo
	Mails []MailDirection `json:"mails" gorm:"many2many:senders"`
	Files []File          `json:"files"`
}

//Client es una estructura que almacena la informacion personal de este cliente asi como su relacion con las ordenes
type Client struct {
	Orm    `json:"orm"`
	Name   string `json:"name"`
	Orders []Order
}

//MailDirection es una estructura que almacena la informacion que un correo electrónico pueda tener
type MailDirection struct {
	Orm       `json:"orm"`
	Direction string `json:"direction"`
}

//File es una estructura que almacena la metadata y localizacion de un archivo
type File struct {
	Orm     `json:"orm"`
	Path    string `json:"path"`
	Title   string `json:"title"`
	OrderID uint   `json:"order_id"`
}

//Profile es una estructura que almacena la informacion de un usuario
type Profile struct {
	Orm                 `json:"orm"`
	DefaultMailConfigID uint         `json:"default_mail_config_id"`
	MailConfig          []MailServer `json:"mail_config"`
	Name                string       `json:"name"`
	ImageAvatarPath     string       `json:"image_avatar_path"`
}

//MailServer almacena la informacion de configuracion de un servidor de correos
type MailServer struct {
	Orm `json:"orm"`
	//Struct fields
	Address string `json:"address"`
	//Alias de este perfil del servidor
	Alias string `json:"alias"`
	//From describe el correo electronico del remitente
	From     string `json:"from"`
	Password string `json:"password"`
	Server   string `json:"server"`
	//Relationship
	ProfiletID uint `json:"profilet_id"`
}
