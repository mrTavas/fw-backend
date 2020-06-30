package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
)

// GetAllCollectors -  Return all workers who is a collector
func GetAllCollectors(c echo.Context) error {

	var OutResponse AllWorkers

	_, err := db.Conn.Query(&OutResponse.Workerss, "SELECT * FROM workers WHERE Collector = ?", true)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
