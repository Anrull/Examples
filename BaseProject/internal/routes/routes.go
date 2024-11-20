package routes

import (
	"Examples/BaseProject/internal/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/api/get_message", handlers.ApiGetMessage)
}
