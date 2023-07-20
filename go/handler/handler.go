package handler

import (
	"net/http"

	"github.com/belarte-tw/copilot-experiments/database"
	"github.com/belarte-tw/copilot-experiments/validate"
	"github.com/labstack/echo/v4"
)

// echo handlers for the endpoints

type Handler struct {
	db *database.DB
}

func NewHandler(db *database.DB) *Handler {
	return &Handler{db: db}
}

// GetActor returns the actor with the given id
func (h *Handler) GetActor(c echo.Context) error {
	id := c.Param("id")
	if !validate.Actor(id) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id: " + id})
	}

	actor, err := h.db.ActorFromId(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "actor not found: " + id})
	}

	return c.JSON(http.StatusOK, actor)
}

func (h *Handler) GetMovie(c echo.Context) error {
	id := c.Param("id")
	if !validate.Movie(id) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id: " + id})
	}

	movie, err := h.db.MovieFromId(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "movie not found: " + id})
	}

	return c.JSON(http.StatusOK, movie)
}
