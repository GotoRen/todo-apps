package controller

import (
	"errors"
	"net/http"

	"github.com/GotoRen/todo-apps/api/internal"
	"github.com/GotoRen/todo-apps/api/logger"
	"github.com/GotoRen/todo-apps/api/model"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func FetchAllTodos(c echo.Context) error {
	todos := make([]model.Todo, 0)

	if err := internal.DBConn().Find(&todos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusOK, model.TodosResponse{})
		}
		logger.LogDebug("Fetched the data", "SELECT", todos)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	return c.JSON(http.StatusOK, todos)
}
