package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// AddManager add information about
func AddManager(c echo.Context) error {

	var inputJSON models.Managers
	var login models.LoginResponse
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	err = db.Conn.Insert(&models.Managers{
		UUID:     uuid.String(),
		Phone:    inputJSON.Phone,
		Initials: inputJSON.Initials,
		Password: inputJSON.Password,
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

	err = login.GenerateJWT(inputJSON)
	if err != nil {
		return err
	}

	return echo.NewHTTPError(http.StatusOK, login)
}
