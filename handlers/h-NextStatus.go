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

// NextStatus Return
func NextStatus(c echo.Context) error {

	var inputJSON nextStatus
	var statuses models.Orders
	var currentStatus string

	var worker models.Workers

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	// Select Worker by id
	err = db.Conn.Model(&worker).Where("ID = ?", inputJSON.NewWorderID).Select()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, "Worker not found. "+err.Error())
	}

	// Select Order by id
	err = db.Conn.Model(&statuses).Where("ID = ?", inputJSON.OrderID).Select()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, "Order not found. "+err.Error())
	}

	// Save order
	SaveOrder(statuses)

	// Get current status
	if statuses.Status.StatusManufacturing == false {

		currentStatus = "manufacturing"
		statuses.Status.StatusManufacturing = true
		statuses.Status.DataManufacturing = time.Now()

	} else if statuses.Status.StatusGrinding == false {

		currentStatus = "grinding"
		statuses.Status.StatusGrinding = true
		statuses.Status.DataGrinding = time.Now()

	} else if statuses.Status.StatusPrinting == false {

		currentStatus = "printing"
		statuses.Status.StatusPrinting = true
		statuses.Status.DataPrinting = time.Now()

	} else {

		currentStatus = "ready"
		statuses.Status.StatusReady = true
		statuses.Status.DataReady = time.Now()

	}

	// Change Worker in current order
	statuses.CurrentWorkerID = worker.ID
	statuses.CurrentWorkerInitials = worker.Initials
	statuses.CurrentWorkerPhone = worker.Phone

	changes := "Заказ переведен на этап \"" + currentStatus + "\". Назначенный работник: \"" + worker.Initials + "\""
	// Save Logs
	err = db.Conn.Insert(&models.OrdersChangesLogs{
		OrderID: statuses.ID,
		Changes: changes,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	// Save current order
	_, err = db.Conn.Model(&statuses).Set("Status = ?, current_worker_id = ?, current_worker_initials = ?, current_worker_phone = ?", statuses.Status, statuses.CurrentWorkerID, statuses.CurrentWorkerInitials, statuses.CurrentWorkerPhone).Where("ID = ?", inputJSON.OrderID).Update()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, currentStatus)

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

		CostManufacturing: order.CostManufacturing,
		CostPainting:      order.CostPainting,
		CostFinishing:     order.CostFinishing,
		CostFull:          order.CostFull,

		Params: order.Params,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "OK")
}
