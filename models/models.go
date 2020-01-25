package models

import "github.com/rs/xid"

//Cliente describe la estructura de un cliente
type Cliente struct {
	Id     xid.ID `json:"id" form:"id" binding:""`
	Nombre string `json:"nombre" form:"nombre" binding:"required"`
}

//Orden describe a una orden de comrpa
type Orden struct {
	Serie    string    `json:"serie"`
	Factura  string    `json:"factura"`
	Archivos []Archivo `json:"archivos"`
}

//Archivos describe a una archivo
type Archivo struct {
	Ruta   string `json:"ruta"`
	Nombre string `json:"nombre"`
}
