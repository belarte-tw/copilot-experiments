package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// start a server with the echo library
// take the address as a parameter
// add the usual middlewares and an inline health check endpoint
func Start(addr string) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	return e.Start(addr)
}
