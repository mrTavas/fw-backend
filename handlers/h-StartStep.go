package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

type nextStatus struct {
	OrderID     int `json:"order_id"`
	NewWorderID int `json:"new_worker_id"`
}

// StartStep - start next step
func StartStep(c echo.Context) error {

	var inputJSON nextStatus
	var order models.Orders
	var currentStatus string
	var worker models.Workers

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	// Select Worker by id
	// id 0 - it's a last step where new worker don't need
	if inputJSON.NewWorderID != 0 {

		err = db.Conn.Model(&worker).Where("ID = ?", inputJSON.NewWorderID).Select()
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, "Worker not found. "+err.Error())
		}
	}

	// Select Order by id
	err = db.Conn.Model(&order).Where("ID = ?", inputJSON.OrderID).Select()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, "Order not found. "+err.Error())
	}

	if order.Status.StatusReady == true {
		return echo.NewHTTPError(http.StatusOK, "Заказ завершен")
	}

	// Get current status
	if order.Status.StatusManufacturingStart == false && order.Status.StatusOfficeEnd == true {

		currentStatus = "manufacturing"
		order.Status.StatusManufacturingStart = true
		order.Status.DataManufacturingStart = time.Now()

	} else if order.Status.StatusGrindingStart == false && order.Status.StatusManufacturingEnd == true {

		currentStatus = "grinding"
		order.Status.StatusGrindingStart = true
		order.Status.DataGrindingStart = time.Now()

	} else if order.Status.StatusPrintingStart == false && order.Status.StatusGrindingEnd == true {

		currentStatus = "printing"
		order.Status.StatusPrintingStart = true
		order.Status.DataPrintingStart = time.Now()

	} else if order.Status.StatusCollectingStart == false && order.Status.StatusPrintingEnd == true {

		currentStatus = "Collecting"
		order.Status.StatusCollectingStart = true
		order.Status.DataCollectingStart = time.Now()

	} else {
		return echo.NewHTTPError(http.StatusOK, "Есть незавершенные этапы.")
	}

	// Change Worker in current order
	order.CurrentWorkerID = worker.ID
	order.CurrentWorkerInitials = worker.Initials
	order.CurrentWorkerPhone = worker.Phone

	changes := "Заказ начал этап \"" + currentStatus + "\". Назначенный на данный этап работник: \"" + worker.Initials + "\""

	// Save Logs
	err = db.Conn.Insert(&models.OrdersChangesLogs{
		OrderID: order.ID,
		Changes: changes,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	// Save current order
	_, err = db.Conn.Model(&order).Set("Status = ?, current_worker_id = ?, current_worker_initials = ?, current_worker_phone = ?", order.Status, order.CurrentWorkerID, order.CurrentWorkerInitials, order.CurrentWorkerPhone).Where("ID = ?", inputJSON.OrderID).Update()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "Начался этап "+currentStatus+".")

}
