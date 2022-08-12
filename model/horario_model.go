package model

type Horario struct {
	Monitor        string `json:"monitor" db:"monitor"`
	DiaDaSemana    int    `json:"dia_da_semana" db:"dia_da_semana"`
	InicioHoras    int    `json:"inicio_horas" db:"inicio_horas"`
	InicioMinutos  int    `json:"inicio_minutos" db:"inicio_minutos"`
	TerminoHoras   int    `json:"termino_horas" db:"termino_horas"`
	TerminoMinutos int    `json:"termino_minutos" db:"termino_minutos"`
}
