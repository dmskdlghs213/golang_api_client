package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func Server(e *echo.Echo) {
	e.GET("/test", TestServer)
}

func TestServer(c echo.Context) error {
	header := c.Request().Header.Get(echo.HeaderAuthorization)
	if header == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "message": "required Authorization token"})
	}

	return c.JSON(http.StatusOK, "success!!")
}
