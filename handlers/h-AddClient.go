package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AddClient add information about Client
// json format:
//{
//	"phone": 89888794747,
//	"initials": "Ivanon I. I.",
//}
func AddClient(c echo.Context) error {

	var inputJSON models.Clients
	rand.Seed(time.Now().UnixNano())

	// var login models.LoginResponse

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	err = db.Conn.Insert(&models.Clients{
		Phone:    inputJSON.Phone,
		Initials: inputJSON.Initials,
		Password: 100000 + rand.Intn(599999),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "OK")
}
