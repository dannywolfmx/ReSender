package models

//Cliente describe la estructura de un cliente
type Cliente struct {
	Id      uint64
	Nombre  string
	Ordenes []Orden
}

//Orden describe a una orden de comrpa
type Orden struct {
	Serie    string
	Factura  string
	Archivos []Archivo
}

//Archivos describe a una archivo
type Archivo struct {
	Ruta   string
	Nombre string
}
