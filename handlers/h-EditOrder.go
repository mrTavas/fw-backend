package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	// local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

type jwtCustomClaims struct {
	UserID int `json:"user_id"`
	Phone  int `json:"login"`
	jwt.StandardClaims
}

// EditOrder -
func EditOrder(c echo.Context) error {

	var inputJSON models.Orders
	var worker models.Workers
	var order models.Orders
	var changes string
	var manager models.Managers
	// var session models.Sessions

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	// Get phone from token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	managerPhone := claims["login"].(float64)

	// Select Manager by id
	err = db.Conn.Model(&manager).Where("Phone = ?", managerPhone).Select()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, "Manager not found. "+err.Error())
	}

	// Select Order by id
	err = db.Conn.Model(&order).Where("ID = ?", inputJSON.ID).Select()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, "Order not found. "+err.Error())
	}

	//***************************
	// Check all changes in order
	//***************************
	if inputJSON.Title != order.Title {

		if inputJSON.Title != "" {
			changes += "Изменено \"Title\" с \"" + order.Title + "\" на \"" + inputJSON.Title + "\". "
		} else {
			inputJSON.Title = order.Title
		}
	}

	if inputJSON.Date != order.Date {

		if inputJSON.Date.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменена \"Дата заказа\" с \"" + order.Date.String() + "\" на \"" + inputJSON.Date.String() + "\". "
		} else {
			inputJSON.Date = order.Date
		}
	}

	if inputJSON.Status.DataOffice != order.Status.DataOffice {

		if inputJSON.Status.DataOffice.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменено \"DataOffice\" с \"" + order.Status.DataOffice.String() + "\" на \"" + inputJSON.Status.DataOffice.String() + "\". "
		} else {
			inputJSON.Status.DataOffice = order.Status.DataOffice
		}
	}

	if inputJSON.Status.DataManufacturing != order.Status.DataManufacturing {

		if inputJSON.Status.DataManufacturing.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменено \"DataManufacturing\" с \"" + order.Status.DataManufacturing.String() + "\" на \"" + inputJSON.Status.DataManufacturing.String() + "\". "
		} else {
			inputJSON.Status.DataManufacturing = order.Status.DataManufacturing
		}

	}

	if inputJSON.Status.DataGrinding != order.Status.DataGrinding {

		if inputJSON.Status.DataGrinding.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменено \"DataGrinding\" с \"" + order.Status.DataGrinding.String() + "\" на \"" + inputJSON.Status.DataGrinding.String() + "\". "

		} else {
			inputJSON.Status.DataGrinding = order.Status.DataGrinding
		}

	}

	if inputJSON.Status.DataPrinting != order.Status.DataReady {

		if inputJSON.Status.DataPrinting.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменено \"DataPrinting\" с \"" + order.Status.DataPrinting.String() + "\" на \"" + inputJSON.Status.DataPrinting.String() + "\". "

		} else {
			inputJSON.Status.DataPrinting = order.Status.DataPrinting
		}

	}

	if inputJSON.Status.DataReady != order.Status.DataReady {

		if inputJSON.Status.DataReady.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменено \"Дата завершения заказа\" с \"" + order.Status.DataReady.String() + "\" на \"" + inputJSON.Status.DataReady.String() + "\". "

		} else {
			inputJSON.Status.DataReady = order.Status.DataReady
		}

	}

	if inputJSON.Status.StatusOffice != order.Status.StatusOffice {
		buff := fmt.Sprintf("Изменено \"StatusOffice\" с \" %v \" на \" %v \". ", order.Status.StatusOffice, inputJSON.Status.StatusOffice)
		changes += buff
	}

	if inputJSON.Status.StatusManufacturing != order.Status.StatusManufacturing {
		buff := fmt.Sprintf("Изменено \"StatusManufacturing\" с \"%v\" на \"%v\". ", order.Status.StatusManufacturing, inputJSON.Status.StatusManufacturing)
		changes += buff
	}

	if inputJSON.Status.StatusGrinding != order.Status.StatusGrinding {
		buff := fmt.Sprintf("Изменено \"StatusGrinding\" с \"%v\" на \"%v\". ", order.Status.StatusGrinding, inputJSON.Status.StatusGrinding)
		changes += buff
	}

	if inputJSON.Status.StatusPrinting != order.Status.StatusPrinting {
		buff := fmt.Sprintf("Изменено \"StatusPrinting\" с \"%v\" на \"%v\". ", order.Status.StatusPrinting, inputJSON.Status.StatusPrinting)
		changes += buff
	}

	if inputJSON.Status.StatusReady != order.Status.StatusReady {
		buff := fmt.Sprintf("Изменено \"Статус завершения заказа\" с \"%v\" на \"%v\". ", order.Status.StatusReady, inputJSON.Status.StatusReady)
		changes += buff
	}

	if inputJSON.ClientID != order.ClientID || inputJSON.ClientInitials != order.ClientInitials || inputJSON.ClientPhone != order.ClientPhone {

		// Try select Client by id
		if inputJSON.ClientID > 0 {
			var client models.Clients

			err = db.Conn.Model(&client).Where("ID = ?", inputJSON.ClientID).Select()
			if err != nil {
				return echo.NewHTTPError(http.StatusOK, "Client not found. "+err.Error())
			}

			// Add Client by id
			inputJSON.ClientInitials = client.Initials
			inputJSON.ClientPhone = client.Phone
		}

		changes += "Изменен \"Клиент заказа\" на \"" + inputJSON.ClientInitials + "\", номер телефона: \"" + strconv.Itoa(inputJSON.ClientPhone) + "\". "

	}

	if inputJSON.CurrentWorkerID != order.CurrentWorkerID {

		// Select Worker by id
		err = db.Conn.Model(&worker).Where("ID = ?", inputJSON.CurrentWorkerID).Select()
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, "Worker not found. "+err.Error())
		}

		// Add worker by id
		inputJSON.CurrentWorkerInitials = worker.Initials
		inputJSON.CurrentWorkerPhone = worker.Phone

		changes += "Изменен \"Текущий работник заказа\" на \"" + inputJSON.CurrentWorkerInitials + "\", номер телефона: \"" + strconv.Itoa(inputJSON.CurrentWorkerPhone) + "\". "

	}

	if inputJSON.Color != order.Color {

		if inputJSON.Color != "" {
			changes += "Изменено \"Color\" с \"" + order.Color + "\" на \"" + inputJSON.Color + "\". "
		} else {
			inputJSON.Color = order.Color
		}
	}

	if inputJSON.Patina != order.Patina {

		if inputJSON.Patina != "" {
			changes += "Изменено \"Patina\" с \"" + order.Patina + "\" на \"" + inputJSON.Patina + "\". "
		} else {
			inputJSON.Patina = order.Patina
		}
	}

	if inputJSON.FasadArticle != order.FasadArticle {

		if inputJSON.FasadArticle != "" {
			changes += "Изменено \"FasadArticle\" с \"" + order.FasadArticle + "\" на \"" + inputJSON.FasadArticle + "\". "
		} else {
			inputJSON.FasadArticle = order.FasadArticle
		}
	}

	if inputJSON.Material != order.Material {

		if inputJSON.Material != "" {
			changes += "Изменено \"Material\" с \"" + order.Material + "\" на \"" + inputJSON.Material + "\". "
		} else {
			inputJSON.Material = order.Material
		}
	}

	if inputJSON.CostManufacturing != order.CostManufacturing {

		if inputJSON.CostManufacturing != 0 {
			changes += "Изменена \"Цена производства\" с \"" + strconv.Itoa(order.CostManufacturing) + "\" на \"" + strconv.Itoa(inputJSON.CostManufacturing) + "\". "
		} else {
			inputJSON.CostManufacturing = order.CostManufacturing
		}
	}

	if inputJSON.CostPainting != order.CostPainting {

		if inputJSON.CostPainting != 0 {
			changes += "Изменена \"Цена покраски\" с \"" + strconv.Itoa(order.CostPainting) + "\" на \"" + strconv.Itoa(inputJSON.CostPainting) + "\". "
		} else {
			inputJSON.CostPainting = order.CostPainting
		}
	}

	if inputJSON.CostFinishing != order.CostFinishing {

		if inputJSON.CostFinishing != 0 {
			changes += "Изменена \"Цена отделки\" с \"" + strconv.Itoa(order.CostFinishing) + "\" на \"" + strconv.Itoa(inputJSON.CostFinishing) + "\". "
		} else {
			inputJSON.CostFinishing = order.CostFinishing
		}
	}

	if inputJSON.CostFull != order.CostFull {

		if inputJSON.CostFull != 0 {
			changes += "Изменена \"Итоговая цена\" с \"" + strconv.Itoa(order.CostFull) + "\" на \"" + strconv.Itoa(inputJSON.CostFull) + "\". "
		} else {
			inputJSON.CostFull = order.CostFull
		}
	}

	// Get chenges in []params
	for i := 0; i < (len(inputJSON.Params)); i++ {

		if inputJSON.Params[i].Title != order.Params[i].Title {

			if inputJSON.Params[i].Title != "" {
				changes += "Изменено \"Комментарий к параметрам\" с \"" + order.Params[i].Title + "\" на \"" + inputJSON.Params[i].Title + "\". "
			} else {
				inputJSON.Params[i].Title = order.Params[i].Title
			}
		}

		if inputJSON.Params[i].Height != order.Params[i].Height {

			if inputJSON.Params[i].Height != 0 {
				changes += "Изменено \"Высота\" с \"" + strconv.Itoa(order.Params[i].Height) + "\" на \"" + strconv.Itoa(inputJSON.Params[i].Height) + "\". "
			} else {
				inputJSON.Params[i].Height = order.Params[i].Height
			}
		}

		if inputJSON.Params[i].Width != order.Params[i].Width {

			if inputJSON.Params[i].Width != 0 {
				changes += "Изменено \"Ширина\" с \"" + strconv.Itoa(order.Params[i].Width) + "\" на \"" + strconv.Itoa(inputJSON.Params[i].Width) + "\". "
			} else {
				inputJSON.Params[i].Width = order.Params[i].Width
			}
		}

		if inputJSON.Params[i].Filenka != order.Params[i].Filenka {

			if inputJSON.Params[i].Filenka != "" {
				changes += "Изменено \"Filenka\" с \"" + order.Params[i].Filenka + "\" на \"" + inputJSON.Params[i].Filenka + "\". "
			} else {
				inputJSON.Params[i].Filenka = order.Params[i].Filenka
			}
		}
	}
	//**********************
	// End check all changes
	//**********************

	// Save Logs
	err = db.Conn.Insert(&models.OrdersChangesLogs{
		OrderID:   inputJSON.ID,
		ManagerID: manager.ID,
		Initials:  manager.Initials,
		Changes:   changes,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	// Save chenged order
	// _, err = db.Conn.Model(&inputJSON).Set("Status = ?", inputJSON).Where("ID = ?", inputJSON.ID).Update()
	err = db.Conn.Update(&inputJSON)

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "OK")

}
