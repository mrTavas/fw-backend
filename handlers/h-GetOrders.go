package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AllOrders - json response type
type AllOrders struct {
	Olders []models.Orders `json:"orders"`
}

// GetOrders Return all managers from db
func GetOrders(c echo.Context) error {

	var OutResponse AllOrders

	_, err := db.Conn.Query(&OutResponse.Olders, "SELECT * FROM orders")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
