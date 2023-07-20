package database_test

import (
	"testing"

	"github.com/belarte-tw/copilot-experiments/database"
	"github.com/stretchr/testify/assert"
)

// test ActorFromId function
// using db.json for test data

func TestActorFromId(t *testing.T) {
	var tests = []struct {
		id    string
		actor database.Actor
	}{
		{"nm0000102", database.Actor{"nm0000102", "Kevin Bacon"}},
		{"nm0000151", database.Actor{"nm0000151", "Morgan Freeman"}},
	}

	db := database.New()

	for _, test := range tests {
		a, err := db.ActorFromId(test.id)
		assert.NoError(t, err)
		assert.Equal(t, test.actor, a)
	}
}

func TestActorNotFound(t *testing.T) {
	db := database.New()
	_, err := db.ActorFromId("nm0000000")
	assert.Error(t, err)
}

// test MovieFromId function
// using db.json for test data

func TestMovieFromId(t *testing.T) {
	var tests = []struct {
		id    string
		movie database.Movie
	}{
		{"tt0109830", database.Movie{"tt0109830", "Forrest Gump"}},
		{"tt0407887", database.Movie{"tt0407887", "The Departed"}},
	}

	db := database.New()

	for _, test := range tests {
		m, err := db.MovieFromId(test.id)
		assert.NoError(t, err)
		assert.Equal(t, test.movie, m)
	}
}

func TestMovieNotFound(t *testing.T) {
	db := database.New()
	_, err := db.MovieFromId("tt0000000")
	assert.Error(t, err)
}

func TestActorsIdsFromMovieId(t *testing.T) {
	var tests = []struct {
		id     string
		actors []string
	}{
		{"tt0104257", []string{"nm0000102", "nm0000129"}},
		{"tt0112384", []string{"nm0000102"}},
	}

	db := database.New()

	for _, test := range tests {
		actors, err := db.ActorsIdsFromMovieId(test.id)
		assert.NoError(t, err)
		assert.Equal(t, test.actors, actors)
	}
}

func TestActorsIdsFromMovieIdNotFound(t *testing.T) {
	db := database.New()
	_, err := db.ActorsIdsFromMovieId("tt0000000")
	assert.Error(t, err)
}

func TestMoviesIdsFromActorId(t *testing.T) {
	var tests = []struct {
		id     string
		movies []string
	}{
		{"nm0000102", []string{"tt0112384", "tt0104257", "tt0087277"}},
		{"nm0000151", []string{}},
	}

	db := database.New()

	for _, test := range tests {
		movies, err := db.MoviesIdsFromActorId(test.id)
		assert.NoError(t, err)
		assert.Equal(t, test.movies, movies)
	}
}

func TestMoviesIdsFromActorIdNotFound(t *testing.T) {
	db := database.New()
	_, err := db.MoviesIdsFromActorId("nm0000000")
	assert.Error(t, err)
}
