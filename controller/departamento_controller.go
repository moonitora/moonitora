package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"net/http"
	"strings"
)

func GetDepartamentos(c *gin.Context) (int, error) {
	type DepartamentoInfo struct {
		Id                    string `json:"id"`
		Nome                  string `json:"nome"`
		MonitoriasAguardando  int    `json:"monitorias_aguardando"`
		MonitoriasConfirmadas int    `json:"monitorias_confirmadas"`
		MonitoriasConcluidas  int    `json:"monitorias_concluidas"`
		MonitoriasCanceladas  int    `json:"monitorias_canceladas"`
		Monitores             int    `json:"monitores"`
	}

	var depts []model.Departamento
	var final []DepartamentoInfo

	if err := repository.DownloadDepartamentos(&depts); err != nil {
		return http.StatusInternalServerError, err
	}

	type CountResult struct {
		Count int `db:"count"`
	}

	db := database.GrabDB()

	for _, dept := range depts {
		var monitorias_aguardando CountResult
		if err := db.Get(&monitorias_aguardando, "SELECT COUNT(*) FROM monitorias WHERE departamento=$1 AND status=$2", dept.Id, 0); err != nil {
			return http.StatusInternalServerError, err
		}

		var monitorias_confirmadas CountResult
		if err := db.Get(&monitorias_confirmadas, "SELECT COUNT(*) FROM monitorias WHERE departamento=$1 AND status=$2", dept.Id, 1); err != nil {
			return http.StatusInternalServerError, err
		}

		var monitorias_concluidas CountResult
		if err := db.Get(&monitorias_concluidas, "SELECT COUNT(*) FROM monitorias WHERE departamento=$1 AND status=$2", dept.Id, 2); err != nil {
			return http.StatusInternalServerError, err
		}

		var monitorias_canceladas CountResult
		if err := db.Get(&monitorias_canceladas, "SELECT COUNT(*) FROM monitorias WHERE departamento=$1 AND status=$2", dept.Id, 3); err != nil {
			return http.StatusInternalServerError, err
		}

		var monitores CountResult
		if err := db.Get(&monitores, "SELECT COUNT(*) FROM usuarios WHERE departamento=$1", dept.Id); err != nil {
			return http.StatusInternalServerError, err
		}

		final = append(final, DepartamentoInfo{
			Id:                    dept.Id,
			Nome:                  dept.Name,
			MonitoriasAguardando:  monitorias_aguardando.Count,
			MonitoriasConfirmadas: monitorias_confirmadas.Count,
			MonitoriasConcluidas:  monitorias_concluidas.Count,
			MonitoriasCanceladas:  monitorias_canceladas.Count,
			Monitores:             monitores.Count,
		})

	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "", "body": final})
	return 0, nil
}

func PostDepartamento(c *gin.Context) (int, error) {
	var departamento model.Departamento
	if err := c.BindJSON(&departamento); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	if departamento.Name == "" {
		return http.StatusBadRequest, errors.New("Especifique um nome para o departamento")
	}

	departamento.Id = strings.ReplaceAll(uuid.New().String(), "-", "")[:10]
	if err := repository.InsertDepartamento(departamento); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Departamento adicionado com sucesso", "body": ""})
	return 0, nil
}

func DeleteDepartamento(c *gin.Context) (int, error) {
	departamento, ok := c.GetQuery("departamento")
	if !ok {
		return http.StatusBadRequest, errors.New("Especifique um departamento")
	}

	if err := repository.DeleteDepartamento(departamento); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Departamento deletado com sucesso", "body": ""})
	return 0, nil
}
