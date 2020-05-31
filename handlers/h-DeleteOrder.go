package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	db "github.com/mrTavas/fw-backend/dbconn"

	//local
	"github.com/mrTavas/fw-backend/models"
)

// OrderID - id of ordr
type OrderID struct {
	ID int `json:"id"`
}

// json format:
//{
//	"id": 8
//}

// DeleteOrder -  delete order by id
func DeleteOrder(c echo.Context) error {

	var inputJSON OrderID

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	order := models.Orders{ID: inputJSON.ID}
	err = db.Conn.Delete(&order)

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "Order deleted")

}
