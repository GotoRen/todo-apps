package controller

import (
	"net/http"
	"time"

	"github.com/GotoRen/todo-apps/api/internal"
	"github.com/GotoRen/todo-apps/api/logger"
	"github.com/GotoRen/todo-apps/api/model"
	"github.com/labstack/echo"
)

func DoneTodo(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, internal.SimpleResponse{Message: "Bad Request(ID is missing)"})
	}

	todo := model.Todo{}
	err := internal.DBConn().Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, internal.SimpleResponse{Message: "DB Error"})
	}

	now := time.Now()
	todo.IsDone = !todo.IsDone
	todo.UpdatedAt = now

	if err = internal.DBConn().Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, internal.SimpleResponse{Message: "DB Error"})
	} else {
		logger.LogDebug("Updated the data", "UPDATE", todo)
	}

	return c.JSON(http.StatusOK, todo)
}
