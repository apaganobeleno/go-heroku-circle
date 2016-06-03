package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apaganobeleno/go-heroku-circle/db"
	"github.com/stretchr/testify/assert"
)

func TestLanding(t *testing.T) {
	DB, _ := db.TestConnection()
	DB.Exec("truncate baq_gophers;")

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	Landing(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
	assert.Contains(t, resp.Body.String(), "GophersBAQ")
	assert.Contains(t, resp.Body.String(), "Aún no tenemos ningun Gopher, nos ayudas?")
}
