package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	db "github.com/mrTavas/fw-backend/dbconn"

	//local
	"github.com/mrTavas/fw-backend/models"
)

//WorkerID - id of worker
type WorkerID struct {
	ID int `json:"id"`
}

// json format:
//{
//	"id": 8
//}

// DeleteWorker -  delete worker by id
func DeleteWorker(c echo.Context) error {

	var inputJSON WorkerID

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	worker := models.Workers{ID: inputJSON.ID}
	err = db.Conn.Delete(&worker)

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "Worker deleted")

}
