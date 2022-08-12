package model

type Departamento struct {
	Id   int    `json:"value" db:"id"`
	Name string `json:"label" db:"title"`
}
