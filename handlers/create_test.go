package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGopher(t *testing.T) {
	rq, _ := http.NewRequest("POST", "/create", nil)
	rq.Form = url.Values{}
	rq.Form.Set("Name", "Antonio")
	rq.Form.Set("Company", "Wawandco")

	rw := httptest.NewRecorder()

	CreateGopher(rw, rq)
	assert.Equal(t, rw.Code, 201)
}
