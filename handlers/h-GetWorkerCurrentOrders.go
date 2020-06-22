package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
)

// GetWorkerCurrentOrders -
func GetWorkerCurrentOrders(c echo.Context) error {

	var OutResponse AllOrders
	var inputJSON WorkerID

	err := c.Bind(&inputJSON)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	err = db.Conn.Model(&OutResponse.Orders).Where("current_worker_id = ?", inputJSON.ID).Select()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
