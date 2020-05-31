package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	db "github.com/mrTavas/fw-backend/dbconn"

	//local
	"github.com/mrTavas/fw-backend/models"
)

// ManagerID - id of manager
type ManagerID struct {
	ID int `json:"id"`
}

// json format:
//{
//	"id": 8
//}

// DeleteManager -  delete manager by id
func DeleteManager(c echo.Context) error {

	var inputJSON ManagerID

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	manager := models.Managers{ID: inputJSON.ID}
	err = db.Conn.Delete(&manager)

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "Manager deleted")

}
