package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	//local
	"github.com/mrTavas/fw-backend/configs"
	"github.com/mrTavas/fw-backend/dbconn"
	h "github.com/mrTavas/fw-backend/handlers"
)

func main() {

	// connecting config struct
	configs.InitConfigs("configs/config")

	// connecting to db
	err := dbconn.Connect()
	if err != nil {
		panic(err)
	}
	defer dbconn.CloseDbConnection(dbconn.Conn)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))

	// Routes
	jwtGroup := e.Group("/api/auth")

	// Clients
	jwtGroup.POST("/NewClient", h.AddClient)
	jwtGroup.GET("/GetClients", h.GetClients)
	jwtGroup.POST("/DeleteClient", h.DeleteClient)

	// PriceList
	jwtGroup.POST("/NewPrice", h.AddPrice)
	jwtGroup.POST("/ChangePrice", h.ChangePrice)
	jwtGroup.POST("/DeletePrice", h.DeletePrice)
	jwtGroup.GET("/GetPriceList", h.GetPriceList)

	// Managers
	jwtGroup.POST("/newmanager", h.AddManager)
	jwtGroup.POST("/DeleteManager", h.DeleteManager)
	jwtGroup.GET("/GetManagers", h.GetManagers)

	// Workers
	jwtGroup.POST("/newworker", h.AddWorker)
	jwtGroup.POST("/DeleteWorker", h.DeleteWorker)
	jwtGroup.GET("/GetWorkers", h.GetWorkers)
	jwtGroup.POST("/GetWorkerCurrentOrders", h.GetWorkerCurrentOrders)
	jwtGroup.POST("/GetWorkerOldOrders", h.GetWorkerOldOrders)

	// Login
	jwtGroup.POST("/login", h.Login)
	jwtGroup.POST("/loginrefresh", h.LoginRefresh)

	// Orders
	jwtGroup.POST("/neworder", h.AddOrder)
	jwtGroup.POST("/DeleteOrder", h.DeleteOrder)
	jwtGroup.GET("/GetOrders", h.GetOrders)
	jwtGroup.GET("/GetSavedOrders", h.GetSavedOrders)
	jwtGroup.POST("/GetOrderStatus", h.GetOrderStatus)
	jwtGroup.POST("/NextStatus", h.NextStatus)
	jwtGroup.POST("/DropStatus", h.DropStatus)

	jwtGroup.POST("/EditOrder", h.EditOrder)

	// Uploads (www/html/uploads)
	jwtGroup.POST("/upload", h.Upload)

	// JWT middleware
	o := e.Group("/api")
	o.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("mySecret"),
	}))

	e.GET("/CreateModels", h.CreateModels)
	e.GET("/DropModels", h.DropModels)
	// e.POST("/CheckPhone", handlers.CheckPhone)
	// e.POST("/Login", handlers.Login)

	o.GET("/main", h.TestJwt)

	// Start server
	e.Logger.Fatal(e.Start(configs.Cfg.Server.MainPort))
}
