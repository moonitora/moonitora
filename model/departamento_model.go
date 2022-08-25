package model

type Departamento struct {
	Id   string `json:"id" db:"uid"`
	Name string `json:"nome" db:"title"`
}
