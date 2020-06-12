package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// DropStatus - drop all statuses in order to default values
func DropStatus(c echo.Context) error {

	var inputJSON OrderID
	var statuses models.Orders

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	err = db.Conn.Model(&statuses).Where("ID = ?", inputJSON.ID).Select()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	// Get current status

	statuses.Status.StatusManufacturing = false
	statuses.Status.StatusGrinding = false
	statuses.Status.StatusPrinting = false
	statuses.Status.StatusReady = false

	statuses.Status.DataManufacturing, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")
	statuses.Status.DataGrinding, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")
	statuses.Status.DataPrinting, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")
	statuses.Status.DataReady, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")

	_, err = db.Conn.Model(&statuses).Set("Status = ?", statuses.Status).Where("ID = ?", inputJSON.ID).Update()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "OK")

}
