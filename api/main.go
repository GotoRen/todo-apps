package main

import (
	"github.com/GotoRen/todo-apps/api/internal"
	"github.com/GotoRen/todo-apps/api/logger"
	"github.com/GotoRen/todo-apps/api/router"
	"github.com/labstack/echo"
)

func main() {
	// Set logger
	logger.InitZap()

	// DB connect
	internal.DBConnecter()
	defer internal.DBClose()

	// DB migration
	internal.DBMigrate()

	// echo API routing
	e := echo.New()
	router.Router(e)

	// HTTP listen
	e.Logger.Fatal(e.Start(":8080"))
}
