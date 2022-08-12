package model

type Horario struct {
	DiaDaSemana    string `json:"dia_da_semana" db:"dia_da_semana"`
	InicioHoras    string `json:"inicio_horas" db:"inicio_horas"`
	InicioMinutos  string `json:"inicio_minutos" db:"inicio_minutos"`
	TerminoHoras   string `json:"termino_horas" db:"termino_horas"`
	TerminoMinutos string `json:"termino_minutos" db:"termino_minutos"`
}
