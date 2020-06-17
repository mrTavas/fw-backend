package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AllClients - json response type
type AllClients struct {
	Clients []models.Clients `json:"clients"`
}

// GetClients Return all clients from db
func GetClients(c echo.Context) error {

	var OutResponse AllClients

	_, err := db.Conn.Query(&OutResponse.Clients, "SELECT * FROM clients")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
