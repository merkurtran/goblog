package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	baseURL := "http://localhost:3002"

	var (
		resp *http.Response
		err  error
	)

	resp, err = http.Get(baseURL + "/")

	assert.NoError(t, err, "Has wrong")
	assert.Equal(t, 200, resp.StatusCode, "Should back 200")
}
