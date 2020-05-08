package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

//TestJwt godoc
func TestJwt(c echo.Context) error {

	return c.String(http.StatusOK, "you're here")
}
