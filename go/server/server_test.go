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

	resp, err := http.Get("http://127.0.0.1:8080/health")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
