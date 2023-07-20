package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/belarte-tw/copilot-experiments/database"
	"github.com/belarte-tw/copilot-experiments/handler"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	db = database.New()
	h  = handler.NewHandler(&db)
)

func TestGetActor(t *testing.T) {
	var tests = map[string]struct {
		id       string
		expected string
		code     int
	}{
		"valid id":      {"nm0000102", `{"imdb_id":"nm0000102","name":"Kevin Bacon"}`, http.StatusOK},
		"also valid id": {"nm0000158", `{"imdb_id":"nm0000158","name":"Tom Hanks"}`, http.StatusOK},
		"unknown id":    {"nm0000000", `{"error":"actor not found: nm0000000"}`, http.StatusNotFound},
		"invalid id":    {"invalid", `{"error":"invalid id: invalid"}`, http.StatusBadRequest},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/actors/"+test.id, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/actors/:id")
			c.SetParamNames("id")
			c.SetParamValues(test.id)

			err := h.GetActor(c)
			assert.NoError(t, err)
			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, test.expected, strings.Trim(rec.Body.String(), " \n\t"))
		})
	}
}

func TestGetMovie(t *testing.T) {
	var tests = map[string]struct {
		id       string
		expected string
		code     int
	}{
		"valid id":      {"tt0109830", `{"imdb_id":"tt0109830","title":"Forrest Gump"}`, http.StatusOK},
		"also valid id": {"tt0112384", `{"imdb_id":"tt0112384","title":"Apollo 13"}`, http.StatusOK},
		"unknown id":    {"tt0000000", `{"error":"movie not found: tt0000000"}`, http.StatusNotFound},
		"invalid id":    {"invalid", `{"error":"invalid id: invalid"}`, http.StatusBadRequest},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/movies/"+test.id, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/movies/:id")
			c.SetParamNames("id")
			c.SetParamValues(test.id)

			err := h.GetMovie(c)
			assert.NoError(t, err)
			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, test.expected, strings.Trim(rec.Body.String(), " \n\t"))
		})
	}
}

func TestActorsFromMoviesId(t *testing.T) {
	// use ActorsIdsFromMovieId to get the actors ids

	var tests = map[string]struct {
		id       string
		expected string
		code     int
	}{
		"valid id":      {"tt0109830", `["Tom Hanks"]`, http.StatusOK},
		"also valid id": {"tt0104257", `["Kevin Bacon","Tom Cruise"]`, http.StatusOK},
		"unknown id":    {"tt0000000", `{"error":"movie not found: tt0000000"}`, http.StatusNotFound},
		"invalid id":    {"invalid", `{"error":"invalid id: invalid"}`, http.StatusBadRequest},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/movies/"+test.id+"/actors", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/movies/:id/actors")
			c.SetParamNames("id")
			c.SetParamValues(test.id)

			err := h.GetActorsFromMovieId(c)
			assert.NoError(t, err)
			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, test.expected, strings.Trim(rec.Body.String(), " \n\t"))
		})
	}
}

func TestMoviesFromActorId(t *testing.T) {
	var tests = map[string]struct {
		id       string
		expected string
		code     int
	}{
		"valid id":      {"nm0000158", `["Forrest Gump","Saving Private Ryan"]`, http.StatusOK},
		"also valid id": {"nm0000102", `["Apollo 13","A Few Good Men","Footloose"]`, http.StatusOK},
		"unknown id":    {"nm0000000", `{"error":"actor not found: nm0000000"}`, http.StatusNotFound},
		"invalid id":    {"invalid", `{"error":"invalid id: invalid"}`, http.StatusBadRequest},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/actors/"+test.id+"/movies", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/actors/:id/movies")
			c.SetParamNames("id")
			c.SetParamValues(test.id)

			err := h.GetMoviesFromActorId(c)
			assert.NoError(t, err)
			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, test.expected, strings.Trim(rec.Body.String(), " \n\t"))
		})
	}
}
