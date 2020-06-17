package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	db "github.com/mrTavas/fw-backend/dbconn"

	//local
	"github.com/mrTavas/fw-backend/models"
)

// ClientID - id of client
type ClientID struct {
	ID int `json:"id"`
}

// json format:
//{
//	"id": 8
//}

// DeleteClient -  delete client by id
func DeleteClient(c echo.Context) error {

	var inputJSON ClientID

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	client := models.Clients{ID: inputJSON.ID}
	err = db.Conn.Delete(&client)

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "Client deleted")

}
