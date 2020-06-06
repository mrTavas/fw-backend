package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AllPriceList -
type AllPriceList struct {
	PriceList []models.PriceList `json:"PriceList"`
}

// GetPriceList -
func GetPriceList(c echo.Context) error {

	var OutResponse AllPriceList

	_, err := db.Conn.Query(&OutResponse.PriceList, "SELECT * FROM price_lists")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, OutResponse)

}
