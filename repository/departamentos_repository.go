package repository

import (
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/model"
)

func DownloadDepartamentos(depts *[]model.Departamento) error {
	db := database.GrabDB()

	if err := db.Select(&depts, "SELECT * FROM departamentos"); err != nil {
		return err
	}

	return nil
}
