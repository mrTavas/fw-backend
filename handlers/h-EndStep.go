package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// EndStep - chanche next end_status to true and inc worker's balance
func EndStep(c echo.Context) error {

	var inputJSON OrderID
	var order models.Orders
	var currentStatus string

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	// Select Order by id
	err = db.Conn.Model(&order).Where("ID = ?", inputJSON.ID).Select()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, "Order not found."+err.Error())
	}

	if order.Status.StatusReady == true {
		return echo.NewHTTPError(http.StatusOK, "Заказ завершен.")
	}

	// Get current status
	if order.Status.StatusOfficeEnd == false {

		currentStatus = "office"
		order.Status.StatusOfficeEnd = true
		order.Status.DataOfficeEnd = time.Now()

	} else if order.Status.StatusManufacturingEnd == false && order.Status.StatusManufacturingStart == true {

		currentStatus = "manufacturing"
		order.Status.StatusManufacturingEnd = true
		order.Status.DataManufacturingEnd = time.Now()

		IncBalance(order.CurrentWorkerID, order.CostCarpenter)

	} else if order.Status.StatusGrindingEnd == false && order.Status.StatusGrindingStart == true {

		currentStatus = "grinding"
		order.Status.StatusGrindingEnd = true
		order.Status.DataGrindingEnd = time.Now()

		IncBalance(order.CurrentWorkerID, order.CostGrinder)

	} else if order.Status.StatusPrintingEnd == false && order.Status.StatusPrintingStart == true {

		currentStatus = "printing"
		order.Status.StatusPrintingEnd = true
		order.Status.DataPrintingEnd = time.Now()

		IncBalance(order.CurrentWorkerID, order.CostPainter)

	} else if order.Status.StatusCollectingEnd == false && order.Status.StatusCollectingStart == true {

		currentStatus = "Collecting"
		order.Status.StatusCollectingEnd = true
		order.Status.DataCollectingEnd = time.Now()

		order.Status.StatusReady = true
		order.Status.DataReady = time.Now()

		IncBalance(order.CurrentWorkerID, order.CostCollector)

	} else {
		return echo.NewHTTPError(http.StatusOK, "Нет начатых этапов.")
	}

	// Changes
	changes := "Заказ завершил этап \"" + currentStatus + "\"."

	// Save Logs
	err = db.Conn.Insert(&models.OrdersChangesLogs{
		OrderID: order.ID,
		Changes: changes,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	// Save order
	if currentStatus != "office" {
		SaveOrder(order)
	}

	// Save current order (nobody has this order)
	_, err = db.Conn.Model(&order).Set("Status = ?, current_worker_id = ?, current_worker_initials = ?, current_worker_phone = ?", order.Status, 0, "", 0).Where("ID = ?", inputJSON.ID).Update()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "Этап "+currentStatus+" завершился.")

}

// IncBalance - increase balance
func IncBalance(workerID int, balance int) error {

	var worker models.Workers

	// Select Worker by id
	err := db.Conn.Model(&worker).Where("ID = ?", workerID).Select()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, "Worker not found. "+err.Error())
	}

	// Set current balance
	_, err = db.Conn.Model(&worker).Set("current_balance = ?", worker.CurrentBalance+balance).Where("ID = ?", workerID).Update()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "OK")
}

// SaveOrder - save order before change status
func SaveOrder(order models.Orders) error {

	err := db.Conn.Insert(&models.SavedOrders{
		OrderID: order.ID,
		Title:   order.Title,
		Date:    order.Date,
		Status:  order.Status,

		ClientID:       order.ClientID,
		ClientInitials: order.ClientInitials,
		ClientPhone:    order.ClientPhone,

		CurrentWorkerID:       order.CurrentWorkerID,
		CurrentWorkerInitials: order.CurrentWorkerInitials,
		CurrentWorkerPhone:    order.CurrentWorkerPhone,

		Color:        order.Color,
		Patina:       order.Patina,
		FasadArticle: order.FasadArticle,
		Material:     order.Material,

		CostCarpenter: order.CostCarpenter,
		CostGrinder:   order.CostGrinder,
		CostPainter:   order.CostPainter,
		CostCollector: order.CostCollector,

		Params: order.Params,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "OK")
}
