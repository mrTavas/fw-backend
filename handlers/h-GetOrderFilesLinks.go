package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// OrdersFilesLinks -
type outResponsePdfandExcle struct {
	Excel string `json:"excel"`
	Pdf   string `json:"pdf"`
}

// GetOrderFilesLinks -
func GetOrderFilesLinks(c echo.Context) error {

	var inputJSON OrderID
	var outResponse models.OrdersFilesLinks

	err := c.Bind(&inputJSON)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	_, err = db.Conn.Query(&outResponse, "SELECT * FROM orders_files_links order by ID DESC LIMIT 1")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var outPdfAndExcel outResponsePdfandExcle
	outPdfAndExcel.Pdf = outResponse.Pdf
	outPdfAndExcel.Excel = outResponse.Excel

	return echo.NewHTTPError(http.StatusOK, outPdfAndExcel)

}
