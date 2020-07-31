package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AllWorkersPost - json response type
type AllWorkersPost struct {
	Сarpenters []models.Workers `json:"carpenters"`
	Grinders   []models.Workers `json:"grinders"`
	Painters   []models.Workers `json:"painters"`
	Collectors []models.Workers `json:"collectors"`
}

// GetWorkersPost -  Return all workers with prof from db
func GetWorkersPost(c echo.Context) error {

	var OutResponse AllWorkersPost

	_, err := db.Conn.Query(&OutResponse.Сarpenters, "SELECT * FROM workers WHERE Сarpenter = ?", true)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	_, err = db.Conn.Query(&OutResponse.Grinders, "SELECT * FROM workers WHERE Grinder = ?", true)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	_, err = db.Conn.Query(&OutResponse.Painters, "SELECT * FROM workers WHERE Painter = ?", true)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	_, err = db.Conn.Query(&OutResponse.Collectors, "SELECT * FROM workers WHERE Collector = ?", true)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
