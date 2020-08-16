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

	Сarpenters []string `json:"carpenters"`
	Grinders   []string `json:"grinders"`
	Painters   []string `json:"painters"`
	Collectors []string `json:"collectors"`
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

			for j := 0; j < (len(orders[i].CurrentWorkers)); j++ {

				OutResponse.Сarpenters[j] = orders[i].CurrentWorkers[j].CurrentWorkerInitials
			}
		}

		if orders[i].Status.StatusGrindingEnd == true && orders[i].Status.StatusPrintingStart == false {

			for j := 0; j < (len(orders[i].CurrentWorkers)); j++ {

				OutResponse.Grinders[j] = orders[i].CurrentWorkers[j].CurrentWorkerInitials
			}
		}

		if orders[i].Status.StatusPrintingEnd == true && orders[i].Status.StatusCollectingStart == false {

			for j := 0; j < (len(orders[i].CurrentWorkers)); j++ {

				OutResponse.Painters[j] = orders[i].CurrentWorkers[j].CurrentWorkerInitials
			}
		}

		if orders[i].Status.StatusCollectingEnd == true {

			for j := 0; j < (len(orders[i].CurrentWorkers)); j++ {

				OutResponse.Collectors[j] = orders[i].CurrentWorkers[j].CurrentWorkerInitials
			}
		}
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
