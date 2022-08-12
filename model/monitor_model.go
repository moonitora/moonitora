package model

type Monitor struct {
	Nome         string `json:"nome" db:"nome"`
	RA           string `json:"ra" db:"ra"`
	Email        string `json:"email" db:"email"`
	Departamento int    `json:"departamento" db:"departamento"`
	Adm          int    `json:"adm" db:"adm"`
}

type MonitorComHorarios struct {
	Monitor  Monitor   `json:"monitor"`
	Horarios []Horario `json:"horarios"`
}

type Login struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
