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
	inputJSON.Status.StatusOfficeStart = true
	inputJSON.Status.DataOfficeStart = time.Now()

	// Add worker by id
	inputJSON.CurrentWorkerInitials = worker.Initials
	inputJSON.CurrentWorkerPhone = worker.Phone

	// Insert
	err = db.Conn.Insert(&models.Orders{
		Date:                  inputJSON.Date,
		Title:                 inputJSON.Title,
		Status:                inputJSON.Status,
		ClientID:              inputJSON.ClientID,
		ClientInitials:        inputJSON.ClientInitials,
		ClientPhone:           inputJSON.ClientPhone,
		CurrentWorkerID:       inputJSON.CurrentWorkerID,
		CurrentWorkerInitials: inputJSON.CurrentWorkerInitials,
		CurrentWorkerPhone:    inputJSON.CurrentWorkerPhone,
		CostCarpenter:         inputJSON.CostCarpenter,
		CostGrinder:           inputJSON.CostGrinder,
		CostPainter:           inputJSON.CostPainter,
		CostCollector:         inputJSON.CostCollector,
		Color:                 inputJSON.Color,
		Patina:                inputJSON.Patina,
		FasadArticle:          inputJSON.FasadArticle,
		Material:              inputJSON.Material,
		Params:                inputJSON.Params,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "Order added")
}
