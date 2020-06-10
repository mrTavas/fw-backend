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

	AddAutoPrice("Ф_01", 4200)
	AddAutoPrice("Ф_02", 4500)
	AddAutoPrice("Ф_03", 5000)
	AddAutoPrice("Ф_04", 5500)
	AddAutoPrice("А_01", 7000)
	AddAutoPrice("А_02", 8000)
	AddAutoPrice("В_01", 7500)
	AddAutoPrice("Прованс", 7000)
	AddAutoPrice("Модерн", 6000)
	AddAutoPrice("Мавзолей", 6500)
	AddAutoPrice("Массив", 8500)

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
