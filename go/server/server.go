package server

import (
	"net/http"

	"github.com/belarte-tw/copilot-experiments/database"
	"github.com/belarte-tw/copilot-experiments/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// start a server with the echo library
// take the address as a parameter
// add the usual middlewares and an inline health check endpoint
func Start(addr string) error {
	db := database.New()
	h := handler.NewHandler(&db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/actors/:id", h.GetActor)
	e.GET("/movies/:id", h.GetMovie)
	return e.Start(addr)
}
