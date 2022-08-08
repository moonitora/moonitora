package model

type Horario struct {
	DiaDaSemana string `json:"dia_da_semana" db:"dia_da_semana"`
	Inicio      string `json:"horario_inicio" db:"horario_termino"`
	Termino     string `json:"horario_termino" db:"horario_termino"`
}
