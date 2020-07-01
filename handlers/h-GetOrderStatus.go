package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// OrderStatusWithWorkers -
type OrderStatusWithWorkers struct {
	OrderStatus models.OrderStatus `json:"statuses"`

	Сarpenter string `json:"carpenter"`
	Grinder   string `json:"grinder"`
	Painter   string `json:"painter"`
	Collector string `json:"collector"`
}

// GetOrderStatus Return
func GetOrderStatus(c echo.Context) error {

	var inputJSON OrderID
	var order models.Orders
	var orders []models.SavedOrders
	var OutResponse OrderStatusWithWorkers

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	err = db.Conn.Model(&order).Where("ID = ?", inputJSON.ID).Select()

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	OutResponse.OrderStatus = order.Status

	_, err = db.Conn.Query(&orders, "SELECT * FROM saved_orders where order_id = ?", inputJSON.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for i := 0; i < (len(orders)); i++ {

		if orders[i].Status.StatusManufacturingEnd == true && orders[i].Status.StatusGrindingStart == false {
			OutResponse.Сarpenter = orders[i].CurrentWorkerInitials
		}

		if orders[i].Status.StatusGrindingEnd == true && orders[i].Status.StatusPrintingStart == false {
			OutResponse.Grinder = orders[i].CurrentWorkerInitials
		}

		if orders[i].Status.StatusPrintingEnd == true && orders[i].Status.StatusCollectingStart == false {
			OutResponse.Painter = orders[i].CurrentWorkerInitials
		}

		if orders[i].Status.StatusCollectingEnd == true {
			OutResponse.Collector = orders[i].CurrentWorkerInitials
		}
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
