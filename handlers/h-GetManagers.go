package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local

	db "github.com/mrTavas/fw-backend/dbconn"
)

type OutJSON struct {
	Respons  string   `json:"response"`
	Managers []string `json:"managers"`
}

// GetManagers sdc
func GetManagers(c echo.Context) error {

	var x []string
	// var Manager models.

	_, err := db.Conn.Query(&x, "SELECT INITIALS FROM managers")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, x)

}
