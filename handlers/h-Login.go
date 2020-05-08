package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local

	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

type inputPass struct {
	Password string `json:"pass"`
	Phone    int    `json:"phone"`
	User     string `json:"user"`
}

// Login This func check login and pass. Gives token.
// json format:
//{
//	"pass" : "qwerty1",
//	"phone": 89888794747,
//	"user": "Client"
//}
func Login(c echo.Context) error {

	var login models.LoginResponse
	var inputJSON inputPass

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	if inputJSON.User == "Manager" {

		var Manager models.Managers
		_, err = db.Conn.Query(&Manager, "SELECT * FROM managers WHERE phone = ?", inputJSON.Phone)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		if inputJSON.Password == Manager.Password {

			err = models.ExpireUserTokens(Manager.UUID)
			if err != nil {
				return err
			}

			err = login.NewRefreshToken(Manager.UUID)
			if err != nil {
				return err
			}

			err = login.GenerateJWT(Manager)
			if err != nil {
				return err
			}

			return echo.NewHTTPError(http.StatusOK, login)
		}
	}

	if inputJSON.User == "Worker" {

		var Worker models.Workers
		_, err = db.Conn.Query(&Worker, "SELECT * FROM workers WHERE phone = ?", inputJSON.Phone)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		if inputJSON.Password == Worker.Password {

			err = models.ExpireUserTokens(Worker.UUID)
			if err != nil {
				return err
			}

			err = login.NewRefreshToken(Worker.UUID)
			if err != nil {
				return err
			}

			err = login.GenerateJWTWorker(Worker)
			if err != nil {
				return err
			}

			return echo.NewHTTPError(http.StatusOK, login)
		}
	}
	return echo.NewHTTPError(http.StatusOK, login)
}
