package controller

import (
	"net/http"

	"github.com/GotoRen/todo-apps/api/internal"
	"github.com/GotoRen/todo-apps/api/logger"
	"github.com/GotoRen/todo-apps/api/model"
	"github.com/labstack/echo"
)

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, internal.SimpleResponse{Message: "Bad Request(ID is missing)"})
	}

	if err := internal.DBConn().Delete(&model.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, internal.SimpleResponse{Message: "DB Error"})
	} else {
		logger.LogDebug("Deleted the data", "DELETE", id)
	}

	return nil
}
