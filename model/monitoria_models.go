package model

type Monitoria struct {
	Monitor  Monitor `json:"monitor" db:"monitor"`
	Conteudo string  `json:"conteudo" db:"conteudo"`
	Dia      int     `json:"dia" db:"dia"`
	Mes      int     `json:"mes" db:"mes"`
	Ano      int     `json:"ano" db:"ano"`
	RAAluno  string  `json:"email_aluno" db:"email_aluno"`
}
