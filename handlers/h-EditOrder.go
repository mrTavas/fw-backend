package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// EditOrder -
func EditOrder(c echo.Context) error {

	var inputJSON models.Orders
	var worker models.Workers

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	// Select Worker by id
	err = db.Conn.Model(&worker).Where("ID = ?", inputJSON.CurrentWorkerID).Select()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, "Worker not found. "+err.Error())
	}

	// Try select Client by id
	if inputJSON.ClientID > 0 {
		var client models.Clients

		err = db.Conn.Model(&client).Where("ID = ?", inputJSON.ClientID).Select()
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, "Client not found. "+err.Error())
		}

		inputJSON.ClientInitials = client.Initials
		inputJSON.ClientPhone = client.Phone
	}

	// Default values
	inputJSON.Status.StatusOffice = true
	inputJSON.Status.DataOffice = time.Now()

	// Add worker by id
	inputJSON.CurrentWorkerInitials = worker.Initials
	inputJSON.CurrentWorkerPhone = worker.Phone

	// _, err = db.Conn.Model(&inputJSON).Set("Status = ?", inputJSON).Where("ID = ?", inputJSON.ID).Update()
	err = db.Conn.Update(&inputJSON)

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "OK")

}
