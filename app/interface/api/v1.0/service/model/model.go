package model

//Profile
type Profile struct {
	//ImageAvatarPath Imagen del perfil
	ImageAvatarPath string `json:"image_avatar_path"`

	//Name nombre del perfil
	Name string `json:"name"`
}

//Client
type Client struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

//MailServer almacena la informacion de configuracion de un servidor de correos
type MailServer struct {
	ID uint `json:"id"`
	//Address represent the server direction including the port.
	//Ej "mail.example.com:25"
	Address string `json:"address"`

	//Alias de este perfil del servidor
	Alias string `json:"alias" validate:"required"`

	//From describe el correo electronico del remitente
	From string `json:"from"`

	//Server ??? //Creo que es server name
	Server string `json:"server"`
}

//File es una estructura que almacena la metadata y localización de un archivo
type File struct {
	ID uint `json:"id"`
	//Title titulo del archivo
	Title string `json:"title"`
}

//MailDirection es una estructura que almacena la informacion que un correo electrónico pueda tener
type MailDirection struct {
	ID uint `json:"id"`
	//Direction direccion de correo electronico. Ej "prueba@gmail.com"
	Direction string `json:"direction" validate:"required,email"`
}

//Order estructura que almacena la informacion de una orden de compra
type Order struct {
	ID uint `json:"id"`
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
