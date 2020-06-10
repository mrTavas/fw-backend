package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AddPrice -
func AddPrice(c echo.Context) error {

	var inputJSON models.PriceList

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	err = db.Conn.Insert(&models.PriceList{
		NAME:  inputJSON.NAME,
		PRICE: inputJSON.PRICE,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "OK")
}
