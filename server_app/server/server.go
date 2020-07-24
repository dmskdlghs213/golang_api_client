package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Server(e *echo.Echo) {
	e.GET("/test", TestServer)
}

func TestServer(c echo.Context) error {

	header := c.Request().Header.Get(echo.HeaderAuthorization)
	if header == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "message": "required Authorization token"})
	}

	log.Info("receive request success")

	return c.JSON(http.StatusOK, "success!!")
}
