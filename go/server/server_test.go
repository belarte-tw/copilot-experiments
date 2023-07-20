package server_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/belarte-tw/copilot-experiments/server"
	"github.com/stretchr/testify/assert"
)

// test the server creation

func TestServerStart(t *testing.T) {
	go server.Start("127.0.0.1:8080")
	time.Sleep(time.Microsecond * 100)

	var tests = []struct {
		url      string
		expected int
	}{
		{"http://127.0.0.1:8080/health", http.StatusOK},
		{"http://127.0.0.1:8080/actors/nm0000102", http.StatusOK},
		{"http://127.0.0.1:8080/movies/tt0109830", http.StatusOK},
		{"http://127.0.0.1:8080/invalid", http.StatusNotFound},
	}

	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			resp, err := http.Get(test.url)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, resp.StatusCode)
		})
	}
}
