package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// ChangePrice -
func ChangePrice(c echo.Context) error {

	var inputJSON models.PriceList

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	_, err = db.Conn.Model(&inputJSON).Set("PRICE = ?", inputJSON.PRICE).Where("NAME = ?", inputJSON.NAME).Update()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "OK")

}
