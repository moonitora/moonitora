package repository

import (
	"database/sql"
	"errors"
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/security"
)

func InsertMonitor(monitor model.Monitor) error {
	db := database.GrabDB()

	var found model.Monitor
	err := db.Get(&found, "SELECT * FROM usuarios WHERE email = $1 OR ra = $2", monitor.Email, monitor.RA)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err != nil {
		return errors.New("Usuário com email ou RA ja registrado.")
	}

	tx := db.MustBegin()
	tx.MustExec(`INSERT INTO usuarios (email, nome, ra, departamento, adm) VALUES ($1,$2,$3,$4,0)`,
		monitor.Email, monitor.Nome, monitor.RA, monitor.Departamento)
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

func DownloadMonitores(departamento int, monitores *[]model.Monitor) error {
	db := database.GrabDB()

	if err := db.Select(monitores, "SELECT * FROM usuarios WHERE departamento=$1", departamento); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("nenhum monitor encontrado")
		}
		return err
	}
	return nil
}

func DownloadMonitor(email string, monitor *model.Monitor) error {
	db := database.GrabDB()

	if err := db.Get(monitor, "SELECT * FROM usuarios WHERE email=$1", email); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("nenhum monitor encontrado")
		}
		return err
	}

	return nil
}

func DownloadMonitorComHorarios(horarios *model.MonitorComHorarios) error {

	return nil
}
