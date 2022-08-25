package model

type Monitoria struct {
	Id           string `json:"id" db:"id"`
	Monitor      string `json:"monitor" db:"monitor"`
	MarcadaPor   string `json:"marcada_por" db:"marcada_por"`
	Departamento string `json:"departamento" db:"departamento"`
	Conteudo     string `json:"conteudo" db:"conteudo"`
	Disciplina   string `json:"disciplina" db:"disciplina"`
	Horario      string `json:"horario" db:"horario"`
	NomeAluno    string `json:"aluno_nome" db:"aluno_nome"`
	RAAluno      string `json:"aluno_ra" db:"aluno_ra"`
	Data         string `json:"data" db:"data"`
	Status       int    `json:"status" db:"status"`
	Sala         string `json:"sala" db:"sala"`
}
