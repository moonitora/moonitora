package model

type Departamento struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"nome" db:"title"`
}
