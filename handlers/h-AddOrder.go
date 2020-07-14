package handlers

import (
	"net/http"

	"time"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AddOrder add new worker in table Workers
func AddOrder(c echo.Context) error {

	var inputJSON models.Orders
	//var outResponse models.Orders

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	// Try Select Worker by id
	if inputJSON.CurrentWorkerID > 0 {
		var worker models.Workers

		err = db.Conn.Model(&worker).Where("ID = ?", inputJSON.CurrentWorkerID).Select()
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, "Worker not found. "+err.Error())
		}

		// Add worker by id
		inputJSON.CurrentWorkerInitials = worker.Initials
		inputJSON.CurrentWorkerPhone = worker.Phone

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
	inputJSON.Status.StatusOfficeStart = true
	inputJSON.Status.DataOfficeStart = time.Now()

	// Insert
	_, err = db.Conn.Model(&inputJSON).Returning("*").Insert()

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, inputJSON.ID)
}
