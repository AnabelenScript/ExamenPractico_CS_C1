package models

type Producto struct {
	ID           string `json:"id"`
	Nombre       string `json:"nombre"`
	Cantidad     int    `json:"cantidad"`
	CodigoBarras string `json:"codigo_barras"`
}
