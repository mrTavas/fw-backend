package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AddWorker add new worker in table Workers
// json format:
//{
//	"phone": 89888794747,
//	"pass" : "qwerty1"
//	"initials": "Ivanon I. I.",
//  "carpenter": false
//  "grinder": false
//  "painter": false
//  "collector": true
//}
func AddWorker(c echo.Context) error {

	var inputJSON models.Workers
	var login models.LoginResponse

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	err = db.Conn.Insert(&models.Workers{
		UUID:      uuid.String(),
		Phone:     inputJSON.Phone,
		Password:  inputJSON.Password,
		Initials:  inputJSON.Initials,
		Сarpenter: inputJSON.Сarpenter,
		Grinder:   inputJSON.Grinder,
		Painter:   inputJSON.Painter,
		Collector: inputJSON.Collector,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	err = models.ExpireUserTokens(uuid.String())
	if err != nil {
		return err
	}

	err = login.NewRefreshToken(uuid.String())
	if err != nil {
		return err
	}

	err = login.GenerateJWTWorker(inputJSON)
	if err != nil {
		return err
	}

	return echo.NewHTTPError(http.StatusOK, login)
}
