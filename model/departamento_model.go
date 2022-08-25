package model

type Departamento struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"nome" db:"title"`
}
