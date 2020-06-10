package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
	//local
)

// PriceName -
type PriceName struct {
	NAME string `json:"name"`
}

// DeletePrice -
func DeletePrice(c echo.Context) error {

	var inputJSON models.PriceList

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	_, err = db.Conn.Model(&inputJSON).Where("NAME = ?", inputJSON.NAME).Delete()

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "Price deleted")

}
