package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AllSavedOrders - json response type
type AllSavedOrders struct {
	SavedOrders []models.SavedOrders `json:"saved_orders"`
}

// GetSavedOrders Return all managers from db
func GetSavedOrders(c echo.Context) error {

	var OutResponse AllSavedOrders

	_, err := db.Conn.Query(&OutResponse.SavedOrders, "SELECT * FROM saved_orders")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
