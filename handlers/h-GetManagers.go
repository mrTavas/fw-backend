package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AllManagers - json response type
type AllManagers struct {
	Managers []models.Managers `json:"managers"`
}

// GetManagers Return all managers from db
func GetManagers(c echo.Context) error {

	var OutResponse AllManagers

	_, err := db.Conn.Query(&OutResponse.Managers, "SELECT * FROM managers")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
