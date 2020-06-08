package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// NextStatus Return
func NextStatus(c echo.Context) error {

	var inputJSON OrderID
	var statuses models.Orders
	var currentStatus string

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	err = db.Conn.Model(&statuses).Where("ID = ?", inputJSON.ID).Select()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	// Get current status
	if statuses.Status.StatusManufacturing == false {

		currentStatus = "manufacturing"
		statuses.Status.StatusManufacturing = true
		statuses.Status.DataManufacturing = time.Now()

		_, err := db.Conn.Model(&statuses).Set("Status = ?", statuses.Status).Where("ID = ?", inputJSON.ID).Update()
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, err.Error())
		}

	} else if statuses.Status.StatusGrinding == false {

		currentStatus = "grinding"
		statuses.Status.StatusGrinding = true
		statuses.Status.DataGrinding = time.Now()

		_, err := db.Conn.Model(&statuses).Set("Status = ?", statuses.Status).Where("ID = ?", inputJSON.ID).Update()
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, err.Error())
		}

	} else if statuses.Status.StatusPrinting == false {

		currentStatus = "printing"
		statuses.Status.StatusPrinting = true
		statuses.Status.DataPrinting = time.Now()

		_, err := db.Conn.Model(&statuses).Set("Status = ?", statuses.Status).Where("ID = ?", inputJSON.ID).Update()
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, err.Error())
		}

	} else {

		currentStatus = "ready"
		statuses.Status.StatusReady = true
		statuses.Status.DataReady = time.Now()

		_, err := db.Conn.Model(&statuses).Set("Status = ?", statuses.Status).Where("ID = ?", inputJSON.ID).Update()
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, err.Error())
		}

	}

	return echo.NewHTTPError(http.StatusOK, currentStatus)

}
