package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	//local
	"github.com/mrTavas/fw-backend/configs"
	"github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/handlers"
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

	// Routes
	jwtGroup := e.Group("/api/auth")
	jwtGroup.POST("/newmanager", h.AddManager)
	jwtGroup.POST("/newworker", h.AddWorker)
	jwtGroup.POST("/login", h.Login)
	jwtGroup.POST("/loginrefresh", h.LoginRefresh)

	//e.GET("/api/GetManagers", h.GetManagers)

	// JWT middleware
	o := e.Group("/api")
	o.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("mySecret"),
	}))

	e.GET("/CreateModels", h.CreateModels)
	e.GET("/DropModels", h.DropModels)
	e.POST("/CheckPhone", handlers.CheckPhone)
	e.POST("/Login", handlers.Login)

	e.GET("/GetManagers", h.GetManagers)
	e.GET("/GetWorkers", h.GetWorkers)

	o.GET("/main", h.TestJwt)

	// Start server
	e.Logger.Fatal(e.Start(configs.Cfg.Server.MainPort))
}
