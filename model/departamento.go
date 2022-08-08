package model

type Curso struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

var Cursos = []Curso{
	Curso{Id: 1, Nome: "Ciências da Natureza e Matemática"},
	Curso{Id: 2, Nome: "Ciências Humanas e Linguagens"},
	Curso{Id: 3, Nome: "Desenvolvimento de Sistemas"},
	Curso{Id: 4, Nome: "Edificações"},
	Curso{Id: 5, Nome: "Geodésia e Cartografia"},
	Curso{Id: 6, Nome: "Enfermagem"},
	Curso{Id: 7, Nome: "Mecânica"},
	Curso{Id: 8, Nome: "Qualidade"},
}

func GetCurso(id int) Curso {
	for _, curso := range Cursos {
		if curso.Id == id {
			return curso
		}
	}
	return Curso{}
}
