package repository

import (
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/model"
)

func DownloadMonitorias(monitor string, monitorias *[]model.Monitoria) error {

	db := database.GrabDB()
	if err := db.Select(monitorias, "SELECT * FROM monitorias WHERE monitor=$1", monitor); err != nil {
		return err
	}

	return nil
}

func DownloadMonitoria(id string, monitoria *model.Monitoria) error {
	db := database.GrabDB()

	if err := db.Get(monitoria, "SELECT * FROM monitorias WHERE id = $1", id); err != nil {
		return err
	}

	return nil
}

func InsertMonitoria(monitoria model.Monitoria) error {
	db := database.GrabDB()
	tx := db.MustBegin()

	tx.MustExec("INSERT INTO monitorias VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", monitoria.Id, monitoria.MarcadaPor, monitoria.Monitor, monitoria.Departamento, monitoria.Conteudo, monitoria.Disciplina, monitoria.Horario, monitoria.NomeAluno, monitoria.RAAluno, monitoria.Data, monitoria.Status)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
