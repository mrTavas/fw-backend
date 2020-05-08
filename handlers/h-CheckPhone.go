package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

type inputPhone struct {
	Phone int    `json:"phone"`
	User  string `json:"user"`
}

// CheckPhone check phone in database
// если нет нужно зарегестрироваться,
// а если есть нужно ввести пароль чтобы войти в свой аккаунт.
// json format:
// {
// 	"phone": 89888794747,
// 	"user": "Client"
// }
func CheckPhone(c echo.Context) error {

	var inputJSON inputPhone
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}
	if inputJSON.User == "Manager" {

		var Client models.Managers
		_, err = db.Conn.Query(&Client, "SELECT * FROM managers WHERE phone = ?", inputJSON.Phone)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return echo.NewHTTPError(http.StatusOK, Client.Phone != inputJSON.Phone)

	} else if inputJSON.User == "Worker" {

		var Worker models.Workers
		_, err = db.Conn.Query(&Worker, "SELECT * FROM workers WHERE phone = ?", inputJSON.Phone)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return echo.NewHTTPError(http.StatusOK, Worker.Phone != inputJSON.Phone)

	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

}
