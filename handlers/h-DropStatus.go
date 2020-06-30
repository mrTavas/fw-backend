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
	statuses.Status.StatusOfficeStart = false

	statuses.Status.StatusManufacturingStart = false
	statuses.Status.StatusManufacturingEnd = false

	statuses.Status.StatusGrindingStart = false
	statuses.Status.StatusGrindingEnd = false

	statuses.Status.StatusPrintingStart = false
	statuses.Status.StatusPrintingEnd = false

	statuses.Status.StatusCollectingStart = false
	statuses.Status.StatusCollectingEnd = false

	statuses.Status.StatusReady = false

	statuses.Status.DataManufacturingStart, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")
	statuses.Status.DataManufacturingEnd, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")

	statuses.Status.DataGrindingStart, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")
	statuses.Status.DataGrindingEnd, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")

	statuses.Status.DataPrintingStart, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")
	statuses.Status.DataPrintingEnd, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")

	statuses.Status.DataCollectingStart, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")
	statuses.Status.DataCollectingEnd, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")

	statuses.Status.DataReady, _ = time.Parse("0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z")

	_, err = db.Conn.Model(&statuses).Set("Status = ?", statuses.Status).Where("ID = ?", inputJSON.ID).Update()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "OK")

}
