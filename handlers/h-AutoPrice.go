package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// GetAutoPrice - add prices in db
func GetAutoPrice() {

	AddAutoPrice("f_01", 4200)
	AddAutoPrice("f_02", 4500)
	AddAutoPrice("f_03", 5000)
	AddAutoPrice("f_04", 5500)
	AddAutoPrice("a_01", 7000)
	AddAutoPrice("a_02", 8000)
	AddAutoPrice("b_01", 7500)
	AddAutoPrice("Provence", 7000)
	AddAutoPrice("Modern", 6000)
	AddAutoPrice("Mausoleum", 6500)
	AddAutoPrice("Massif", 8500)

}

// AddAutoPrice - auto price
func AddAutoPrice(name string, price int) error {

	err := db.Conn.Insert(&models.PriceList{
		NAME:  name,
		PRICE: price,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "OK")
}
