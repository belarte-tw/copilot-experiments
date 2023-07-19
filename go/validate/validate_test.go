package validate_test

import (
	"testing"

	"github.com/belarte-tw/copilot-experiments/validate"
	"github.com/stretchr/testify/assert"
)

// validate is a package to validate imdb ids for actors and movies

func TestValidateActor(t *testing.T) {
	var tests = []struct {
		id    string
		valid bool
	}{
		{"nm0000001", true},
		{"nm0000002", true},
		{"nm000000", false},
		{"nm00000000", false},
		{"nm000000a", false},
		{"nm0000001a", false},
		{"tt0000001", false},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, validate.Actor(test.id))
	}
}

func TestValidateMovie(t *testing.T) {
	var tests = []struct {
		id    string
		valid bool
	}{
		{"tt0000001", true},
		{"tt0000002", true},
		{"tt000000", false},
		{"tt00000000", false},
		{"tt000000a", false},
		{"tt0000001a", false},
		{"nm0000001", false},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, validate.Movie(test.id))
	}
}
