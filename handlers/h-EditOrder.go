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
	var currentWorkersForLogs string

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

	if inputJSON.Status.DataOfficeStart != order.Status.DataOfficeStart {

		if inputJSON.Status.DataOfficeStart.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменено \"DataOfficeStart\" с \"" + order.Status.DataOfficeStart.String() + "\" на \"" + inputJSON.Status.DataOfficeStart.String() + "\". "
		} else {
			inputJSON.Status.DataOfficeStart = order.Status.DataOfficeStart
		}
	}

	if inputJSON.Status.DataManufacturingStart != order.Status.DataManufacturingStart {

		if inputJSON.Status.DataManufacturingStart.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменено \"DataManufacturingStart\" с \"" + order.Status.DataManufacturingStart.String() + "\" на \"" + inputJSON.Status.DataManufacturingStart.String() + "\". "
		} else {
			inputJSON.Status.DataManufacturingStart = order.Status.DataManufacturingStart
		}

	}

	if inputJSON.Status.DataGrindingStart != order.Status.DataGrindingStart {

		if inputJSON.Status.DataGrindingStart.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменено \"DataGrindingStart\" с \"" + order.Status.DataGrindingStart.String() + "\" на \"" + inputJSON.Status.DataGrindingStart.String() + "\". "

		} else {
			inputJSON.Status.DataGrindingStart = order.Status.DataGrindingStart
		}

	}

	if inputJSON.Status.DataPrintingStart != order.Status.DataPrintingStart {

		if inputJSON.Status.DataPrintingStart.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменено \"DataPrintingStart\" с \"" + order.Status.DataPrintingStart.String() + "\" на \"" + inputJSON.Status.DataPrintingStart.String() + "\". "

		} else {
			inputJSON.Status.DataPrintingStart = order.Status.DataPrintingStart
		}

	}

	if inputJSON.Status.DataCollectingStart != order.Status.DataCollectingStart {

		if inputJSON.Status.DataCollectingStart.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменено \"DataCollectingStart\" с \"" + order.Status.DataCollectingStart.String() + "\" на \"" + inputJSON.Status.DataCollectingStart.String() + "\". "

		} else {
			inputJSON.Status.DataCollectingStart = order.Status.DataCollectingStart
		}

	}

	if inputJSON.Status.DataReady != order.Status.DataReady {

		if inputJSON.Status.DataReady.String() != "0001-01-01 00:00:00 +0000 UTC" {
			changes += "Изменено \"Дата завершения заказа\" с \"" + order.Status.DataReady.String() + "\" на \"" + inputJSON.Status.DataReady.String() + "\". "

		} else {
			inputJSON.Status.DataReady = order.Status.DataReady
		}

	}

	if inputJSON.Status.StatusOfficeStart != order.Status.StatusOfficeStart {
		buff := fmt.Sprintf("Изменено \"StatusOfficeStart\" с \" %v \" на \" %v \". ", order.Status.StatusOfficeStart, inputJSON.Status.StatusOfficeStart)
		changes += buff
	}

	if inputJSON.Status.StatusManufacturingStart != order.Status.StatusManufacturingStart {
		buff := fmt.Sprintf("Изменено \"StatusManufacturingStart\" с \"%v\" на \"%v\". ", order.Status.StatusManufacturingStart, inputJSON.Status.StatusManufacturingStart)
		changes += buff
	}

	if inputJSON.Status.StatusGrindingStart != order.Status.StatusGrindingStart {
		buff := fmt.Sprintf("Изменено \"StatusGrindingStart\" с \"%v\" на \"%v\". ", order.Status.StatusGrindingStart, inputJSON.Status.StatusGrindingStart)
		changes += buff
	}

	if inputJSON.Status.StatusPrintingStart != order.Status.StatusPrintingStart {
		buff := fmt.Sprintf("Изменено \"StatusPrintingStart\" с \"%v\" на \"%v\". ", order.Status.StatusPrintingStart, inputJSON.Status.StatusPrintingStart)
		changes += buff
	}

	if inputJSON.Status.StatusCollectingStart != order.Status.StatusCollectingStart {
		buff := fmt.Sprintf("Изменено \"StatusCollectingStart\" с \"%v\" на \"%v\". ", order.Status.StatusCollectingStart, inputJSON.Status.StatusCollectingStart)
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

	for i := 0; i < (len(inputJSON.CurrentWorkers)); i++ {

		if len(order.CurrentWorkers) != 0 {

			if inputJSON.CurrentWorkers[i].CurrentWorkerID != order.CurrentWorkers[i].CurrentWorkerID {

				// Select Worker by id
				err = db.Conn.Model(&worker).Where("ID = ?", inputJSON.CurrentWorkers[i].CurrentWorkerID).Select()
				if err != nil {
					return echo.NewHTTPError(http.StatusOK, "Worker not found. "+err.Error())
				}

				// Add worker by id
				inputJSON.CurrentWorkers[i].CurrentWorkerInitials = worker.Initials
				inputJSON.CurrentWorkers[i].CurrentWorkerPhone = worker.Phone

				currentWorkersForLogs += worker.Initials + " "
			}

		} else {

			// Select Worker by id
			err = db.Conn.Model(&worker).Where("ID = ?", inputJSON.CurrentWorkers[i].CurrentWorkerID).Select()
			if err != nil {
				return echo.NewHTTPError(http.StatusOK, "Worker not found. "+err.Error())
			}

			// Add worker by id
			inputJSON.CurrentWorkers[i].CurrentWorkerInitials = worker.Initials
			inputJSON.CurrentWorkers[i].CurrentWorkerPhone = worker.Phone

			currentWorkersForLogs += worker.Initials + " "

		}
	}

	if currentWorkersForLogs != "" {

		changes += "Измененены \"текущие работники заказа\" на: \"" + currentWorkersForLogs + ". "
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

	if inputJSON.CostCarpenter != order.CostCarpenter {

		if inputJSON.CostCarpenter != 0 {
			changes += "Изменена \"Цена столярных работ\" с \"" + strconv.Itoa(order.CostCarpenter) + "\" на \"" + strconv.Itoa(inputJSON.CostCarpenter) + "\". "
		} else {
			inputJSON.CostCarpenter = order.CostCarpenter
		}
	}

	if inputJSON.CostGrinder != order.CostGrinder {

		if inputJSON.CostGrinder != 0 {
			changes += "Изменена \"Цена шлифовки\" с \"" + strconv.Itoa(order.CostGrinder) + "\" на \"" + strconv.Itoa(inputJSON.CostGrinder) + "\". "
		} else {
			inputJSON.CostGrinder = order.CostGrinder
		}
	}

	if inputJSON.CostPainter != order.CostPainter {

		if inputJSON.CostPainter != 0 {
			changes += "Изменена \"Цена покраски\" с \"" + strconv.Itoa(order.CostPainter) + "\" на \"" + strconv.Itoa(inputJSON.CostPainter) + "\". "
		} else {
			inputJSON.CostPainter = order.CostPainter
		}
	}

	if inputJSON.CostCollector != order.CostCollector {

		if inputJSON.CostCollector != 0 {
			changes += "Изменена \"Цена сборки\" с \"" + strconv.Itoa(order.CostCollector) + "\" на \"" + strconv.Itoa(inputJSON.CostCollector) + "\". "
		} else {
			inputJSON.CostCollector = order.CostCollector
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
