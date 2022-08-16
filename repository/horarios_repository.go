package repository

import (
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/model"
)

func DownloadHorarios(monitor string, horarios *[]model.Horario) error {
	db := database.GrabDB()

	if err := db.Select(horarios, "SELECT * FROM horarios WHERE monitor=$1", monitor); err != nil {
		return err
	}

	return nil
}

func DownloadHorario(id string, horario *model.Horario) error {
	db := database.GrabDB()

	if err := db.Get(horario, "SELECT * FROM horarios WHERE id=$1", id); err != nil {
		return err
	}

	return nil
}

func InsertHorario(horario model.Horario) error {
	db := database.GrabDB()
	tx := db.MustBegin()

	tx.MustExec("INSERT INTO horarios VALUES ($1,$2,$3,$4,$5,$6,$7)", horario.Id, horario.Monitor, horario.DiaDaSemana, horario.InicioHoras, horario.InicioMinutos, horario.TerminoHoras, horario.TerminoMinutos)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func DeleteHorario(horario string) error {
	db := database.GrabDB()
	tx := db.MustBegin()

	tx.MustExec("DELETE FROM horarios WHERE id=$1", horario)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
