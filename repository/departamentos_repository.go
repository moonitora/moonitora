package repository

import (
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/model"
)

func DownloadDepartamentos(depts *[]model.Departamento) error {
	db := database.GrabDB()

	if err := db.Select(depts, "SELECT * FROM departamentos"); err != nil {
		return err
	}

	return nil
}

func InsertDepartamento(departamento model.Departamento) error {
	db := database.GrabDB()
	tx := db.MustBegin()

	tx.MustExec("INSERT INTO departamentos VALUES ($1,$2)", departamento.Id, departamento.Name)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func DeleteDepartamento(id string) error {
	db := database.GrabDB()

	tx := db.MustBegin()
	tx.MustExec("DELETE FROM monitorias WHERE departamento=$1", id)
	if err := tx.Commit(); err != nil {
		return err
	}

	var monitores []model.Monitor
	if err := DownloadMonitores(id, &monitores); err != nil {
		return err
	}

	for _, monitor := range monitores {
		tx = db.MustBegin()
		tx.MustExec("DELETE FROM horarios WHERE monitor=$1", monitor.Email)
		if err := tx.Commit(); err != nil {
			return err
		}
	}

	tx = db.MustBegin()
	tx.MustExec("DELETE FROM departamentos WHERE id=$1", id)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
