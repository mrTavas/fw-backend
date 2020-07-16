package handlers

import (
	"net/http"

	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"

	//local
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

//DropModels Delete all models in database.
func DropModels(c echo.Context) error {

	for _, model := range []interface{}{
		&models.OrdersFilesLinks{},
		&models.OrdersChangesLogs{},
		&models.SavedOrders{},
		&models.Clients{},
		&models.PriceList{},
		&models.Orders{},
		&models.Workers{},
		&models.Managers{},
		&models.Sessions{}} {
		err := db.Conn.DropTable(model, &orm.DropTableOptions{})
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, err.Error())
			//panic(err)
		}
	}

	return c.String(http.StatusOK, "Models Deleted/Dropped")
}
