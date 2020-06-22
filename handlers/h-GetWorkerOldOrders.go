package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
)

// GetWorkerOldOrders -
func GetWorkerOldOrders(c echo.Context) error {

	var OutResponse AllSavedOrders
	var inputJSON WorkerID

	err := c.Bind(&inputJSON)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	err = db.Conn.Model(&OutResponse.SavedOrders).Where("current_worker_id = ?", inputJSON.ID).Select()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
