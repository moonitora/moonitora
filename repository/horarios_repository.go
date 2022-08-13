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
