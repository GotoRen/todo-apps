package router

import (
	"net/http"
	"os"

	"github.com/GotoRen/todo-apps/api/internal"
	"github.com/labstack/echo"
)

func healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, internal.SimpleResponse{Message: "Health check endpoint"})
}

func check(c echo.Context) error {
	return c.JSON(http.StatusOK, internal.SimpleResponse{Message: os.Getenv("HOST_NAME")})
}
