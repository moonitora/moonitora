package repository

import (
	"database/sql"
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/security"
)

func InsertMonitor(monitor model.Monitor) error {
	db := database.GrabDB()

	var found model.Monitor
	err := db.Get(&found, "SELECT * FROM usuarios WHERE email = $1", monitor.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	tx := db.MustBegin()
	tx.MustExec(`INSERT INTO usuarios (email, nome, ra, departamento, curso) VALUES ($1,$2,$3,$4,$5)`,
		monitor.Email, monitor.Nome, monitor.RA, monitor.Departamento, monitor.Curso)
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func InsertLogin(login model.Login) error {
	db := database.GrabDB()
	crypt, _ := security.Hash(login.Password)
	tx2 := db.MustBegin()
	tx2.MustExec("INSERT INTO login VALUES ($1,$2)", login.Email, string(crypt))
	if err := tx2.Commit(); err != nil {
		return err
	}
	return nil
}
