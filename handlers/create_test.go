package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/apaganobeleno/go-heroku-circle/db"
	"github.com/apaganobeleno/go-heroku-circle/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateGopher(t *testing.T) {
	DB, _ := db.TestConnection()
	DB.Exec("truncate baq_gophers;")

	rq, _ := http.NewRequest("POST", "/create", nil)
	rq.Form = url.Values{}
	rq.Form.Set("Name", "Antonio")
	rq.Form.Set("Company", "Wawandco")

	rw := httptest.NewRecorder()
	CreateGopher(rw, rq)

	var count int
	DB.Model(&models.Gopher{}).Count(&count)

	assert.Equal(t, count, 1)
	assert.Equal(t, rw.Code, 302)
}
