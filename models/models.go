package models

//Cliente describe la estructura de un cliente
type Cliente struct {
	Id     uint64 `json:"id" form:"id" binding:"required"`
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
