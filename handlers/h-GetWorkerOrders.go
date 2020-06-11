package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
)

// WorkerPhone -
type WorkerPhone struct {
	Phone int `json:"current_worker_phone"`
}

// GetWorkerOrders -
func GetWorkerOrders(c echo.Context) error {

	var OutResponse AllOrders
	var inputJSON WorkerPhone

	err := c.Bind(&inputJSON)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	err = db.Conn.Model(&OutResponse.Orders).Where("current_worker_phone = ?", inputJSON.Phone).Select()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
