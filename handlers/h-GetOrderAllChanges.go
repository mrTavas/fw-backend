package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AllOrderChanges - json response type
type AllOrderChanges struct {
	Changes []models.OrdersChangesLogs `json:"changes"`
}

// GetOrderAllChanges -
func GetOrderAllChanges(c echo.Context) error {

	var inputJSON OrderID
	var OutResponse AllOrderChanges

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	err = db.Conn.Model(&OutResponse.Changes).Where("order_id = ?", inputJSON.ID).Select()

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	for i := 0; i < len(OutResponse.Changes); i++ {
		OutResponse.Changes[i].Changes = RemoveQuotes(OutResponse.Changes[i].Changes)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
