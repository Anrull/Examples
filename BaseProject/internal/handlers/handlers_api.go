package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ApiGetMessage(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
