package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AddOrder add new worker in table Workers
// json format:
// {
//  "status": "Office"
// 	"client_initials": "Clientov A.V.",
// 	"client_phone" : 79888563211,
// 	"current_worker_initials": "Ivanon I. I.",
//  "current_worker_phone": 7988121212,
//  "cost_manufacturing": 3000,
//  "cost_painting": 2000,
//  "cost_finishing": 1500,
//  "cost_full": 7500
// }
func AddOrder(c echo.Context) error {

	var inputJSON models.Orders

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	err = db.Conn.Insert(&models.Orders{
		Date:                  inputJSON.Date,
		Status:                inputJSON.Status,
		ClientInitials:        inputJSON.ClientInitials,
		ClientPhone:           inputJSON.ClientPhone,
		CurrentWorkerInitials: inputJSON.CurrentWorkerInitials,
		CurrentWorkerPhone:    inputJSON.CurrentWorkerPhone,
		CostManufacturing:     inputJSON.CostManufacturing,
		CostPainting:          inputJSON.CostPainting,
		CostFinishing:         inputJSON.CostFinishing,
		CostFull:              inputJSON.CostFull,
		Params:                inputJSON.Params,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "Order added")
}
