package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/authorization"
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"net/http"
)

type IncomingUser struct {
	Monitor model.Monitor `json:"monitor"`
	Login   model.Login   `json:"login"`
}

func FetchMonitores(c *gin.Context) (int, error) {
	dept, ok := c.GetQuery("departamento")
	_, okDeep := c.GetQuery("deep")
	if !ok {
		return http.StatusBadRequest, errors.New("especifique um departamento")
	}

	type MonitorInfo struct {
		Monitor               model.Monitor `json:"monitor"`
		MonitoriasAguardando  int           `json:"monitorias_aguardando"`
		MonitoriasConfirmadas int           `json:"monitorias_confirmadas"`
		MonitoriasConcluidas  int           `json:"monitorias_concluidas"`
		MonitoriasCanceladas  int           `json:"monitorias_canceladas"`
	}

	var final []MonitorInfo

	db := database.GrabDB()

	type CountResult struct {
		Count int `db:"count"`
	}

	var monitores []model.Monitor
	if err := repository.DownloadMonitores(dept, &monitores); err != nil {
		return http.StatusInternalServerError, err
	}
	for _, monitor := range monitores {

		var (
			monitorias_aguardando  CountResult
			monitorias_confirmadas CountResult
			monitorias_concluidas  CountResult
			monitorias_canceladas  CountResult
		)

		if okDeep {
			if err := db.Get(&monitorias_aguardando, "SELECT COUNT(*) FROM monitorias WHERE monitor=$1 AND status=$2", monitor.Email, 0); err != nil {
				return http.StatusInternalServerError, err
			}

			if err := db.Get(&monitorias_confirmadas, "SELECT COUNT(*) FROM monitorias WHERE monitor=$1 AND status=$2", monitor.Email, 1); err != nil {
				return http.StatusInternalServerError, err
			}

			if err := db.Get(&monitorias_concluidas, "SELECT COUNT(*) FROM monitorias WHERE monitor=$1 AND status=$2", monitor.Email, 2); err != nil {
				return http.StatusInternalServerError, err
			}

			if err := db.Get(&monitorias_canceladas, "SELECT COUNT(*) FROM monitorias WHERE monitor=$1 AND status=$2", monitor.Email, 3); err != nil {
				return http.StatusInternalServerError, err
			}
		}

		final = append(final, MonitorInfo{
			Monitor:               monitor,
			MonitoriasAguardando:  monitorias_aguardando.Count,
			MonitoriasConfirmadas: monitorias_confirmadas.Count,
			MonitoriasConcluidas:  monitorias_concluidas.Count,
			MonitoriasCanceladas:  monitorias_canceladas.Count,
		})

	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "", "body": final})
	return 0, nil
}

func Register(c *gin.Context) (int, error) {
	incoming := IncomingUser{}
	if err := c.BindJSON(&incoming); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	if err := repository.InsertMonitor(incoming.Monitor); err != nil {
		return http.StatusBadRequest, err
	}

	if err := repository.InsertLogin(incoming.Login); err != nil {
		return http.StatusBadRequest, err
	}

	token := authorization.GenerateToken(incoming.Monitor.Email)

	c.JSON(http.StatusOK, gin.H{"jwt": token, "body": incoming.Monitor, "status": true, "message": "Usuário cadastrado com sucesso"})
	return 0, nil
}
