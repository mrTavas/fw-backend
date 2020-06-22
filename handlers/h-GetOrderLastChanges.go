package handlers

import (
	"bytes"
	"net/http"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// GetOrderLastChanges -
func GetOrderLastChanges(c echo.Context) error {

	var inputJSON OrderID
	var OutResponse []models.OrdersChangesLogs

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	err = db.Conn.Model(&OutResponse).Where("order_id = ?", inputJSON.ID).Select()

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	if len(OutResponse) == 0 {
		return echo.NewHTTPError(http.StatusOK, "Нет изменений")
	}

	OutResponse[len(OutResponse)-1].Changes = RemoveQuotes(OutResponse[len(OutResponse)-1].Changes)
	return echo.NewHTTPError(http.StatusOK, OutResponse[len(OutResponse)-1])

}

// RemoveQuotes -sw
func RemoveQuotes(s string) string {
	var b bytes.Buffer
	for _, r := range s {
		if r != '"' && r != '\'' {
			b.WriteRune(r)
		}
	}

	return b.String()
}
