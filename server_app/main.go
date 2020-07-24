package main

import (
	"net/http"

	"github.com/dmskdlghs213/golang_api_client_and_server/server_app/server"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	server.Server(e)

	e.Logger.Fatal(e.Start(":8080")) // listen and serve on :8080
}
