package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
)

// GetWorkersPost -  Return all managers from db
func GetWorkersPost(c echo.Context) error {

	var OutResponse AllWorkers

	_, err := db.Conn.Query(&OutResponse.Workerss, "SELECT * FROM workers order by ID")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
