package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
)

// GetAllCarpenters -  Return all workers who is a carpenter
func GetAllCarpenters(c echo.Context) error {

	var OutResponse AllWorkers

	_, err := db.Conn.Query(&OutResponse.Workerss, "SELECT * FROM workers WHERE Сarpenter = ?", true)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
