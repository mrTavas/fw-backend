package handlers

import (
	"net/http"

	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

//CreateModels Create all models in database
func CreateModels(c echo.Context) error {

	for _, model := range []interface{}{
		&models.OrdersChangesLogs{},
		&models.SavedOrders{},
		&models.Clients{},
		&models.PriceList{},
		&models.Workers{},
		&models.Managers{},
		&models.Orders{},
		&models.Sessions{}} {
		err := db.Conn.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusOK, err.Error())
			//panic(err)
		}
	}

	GetAutoPrice()
	return c.String(http.StatusOK, "Models Created")
}
