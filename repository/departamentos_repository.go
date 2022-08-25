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
