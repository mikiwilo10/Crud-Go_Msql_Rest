package model

type Persona struct {
	IdPersona int    `json:"IdPersona"`
	Nombre    string `json:"Nombre"`
	Apellido  string `json:"Apellido"`
	Email     string `json:"Email"`
	Genero    string `json:"Genero"`
}
type AllPersonas []Persona
