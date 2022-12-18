package controller

import (
	"net/http"

	"github.com/GotoRen/todo-apps/api/internal"
	"github.com/GotoRen/todo-apps/api/logger"
	"github.com/GotoRen/todo-apps/api/model"
	"github.com/labstack/echo"
)

func PostTodo(c echo.Context) error {
	todo := model.Todo{}

	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, internal.SimpleResponse{Message: "Bad Request"})
	}

	if err := c.Validate(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, internal.SimpleResponse{Message: "Bad Request(Validation error)"})
	}

	err := internal.DBConn().Create(&todo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, internal.SimpleResponse{Message: "DB Error"})
	} else {
		logger.LogDebug("Inserted the data", "INSERT", todo)
	}

	return nil
}
