package model

type Monitor struct {
	Nome         string    `json:"nome" db:"nome"`
	RA           string    `json:"ra" db:"ra"`
	Email        string    `json:"email" db:"email"`
	Departamento int       `json:"departamento" db:"departamento"`
	Horarios     []Horario `json:"horarios" db:"horarios"`
	Curso        int       `json:"curso" db:"curso"`
}

type Login struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
