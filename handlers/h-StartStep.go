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
	OrderID      int        `json:"order_id"`
	NewWorkersID []WorkerID `json:"new_workers_id"`
}

// StartStep - start next step
func StartStep(c echo.Context) error {

	var inputJSON nextStatus
	var order models.Orders
	var currentStatus string
	var worker models.Workers

	var currentWorker models.CurrentWorker
	var workers []models.CurrentWorker

	var currentWorkersForLogs string

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	for i := 0; i < (len(inputJSON.NewWorkersID)); i++ {

		// Select Worker by id
		// id 0 - it's a last step where new worker don't need
		if inputJSON.NewWorkersID[i].ID != 0 {

			err = db.Conn.Model(&worker).Where("ID = ?", inputJSON.NewWorkersID[i].ID).Select()
			if err != nil {
				return echo.NewHTTPError(http.StatusOK, "Worker not found. "+err.Error())
			}

			currentWorker.CurrentWorkerID = worker.ID
			currentWorker.CurrentWorkerInitials = worker.Initials
			currentWorker.CurrentWorkerPhone = worker.Phone

			workers = append(workers, currentWorker)

			currentWorkersForLogs += worker.Initials + " "

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

	// // Change Worker in current order
	// for i := 0; i < (len(inputJSON.NewWorkersID)); i++ {
	// 	println("pass -> +1")

	// 	order.CurrentWorkers[i].CurrentWorkerID = workers[i].CurrentWorkerID
	// 	order.CurrentWorkers[i].CurrentWorkerInitials = workers[i].CurrentWorkerInitials
	// 	order.CurrentWorkers[i].CurrentWorkerPhone = workers[i].CurrentWorkerPhone

	// 	// currentWorkersForLogs += workers[i].CurrentWorkerInitials + " "

	// 	println("pass -> +2")
	// }

	changes := "Заказ начал этап \"" + currentStatus + "\". Назначенные на данный этап работники: \"" + currentWorkersForLogs + "\""

	// Save Logs
	err = db.Conn.Insert(&models.OrdersChangesLogs{
		OrderID: order.ID,
		Changes: changes,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	// Save current order
	_, err = db.Conn.Model(&order).Set("Status = ?, current_workers = ?", order.Status, workers).Where("ID = ?", inputJSON.OrderID).Update()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "Начался этап "+currentStatus+".")

}
