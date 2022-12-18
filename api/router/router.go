package router

import (
	"net/http"

	"github.com/GotoRen/todo-apps/api/controller"
	"github.com/GotoRen/todo-apps/api/internal"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Router(e *echo.Echo) {
	// CORS (Cross-Origin Resource Sharing) measures
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Validator = &internal.CustomValidator{
		Validator: validator.New(),
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthz", healthz)

	api := e.Group("/api")
	api.GET("/check", check)
	api.GET("/todos", controller.FetchAllTodos)
	api.POST("/todo", controller.PostTodo)
	api.POST("/todo/:id/done", controller.DoneTodo)
	api.DELETE("/todo/:id", controller.DeleteTodo)
}
