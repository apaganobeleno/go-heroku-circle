package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apaganobeleno/go-heroku-circle/db"
	"github.com/apaganobeleno/go-heroku-circle/models"
	"github.com/stretchr/testify/assert"
	"github.com/wawandco/fako"
)

func TestLanding(t *testing.T) {
	DB, _ := db.TestConnection()
	DB.Exec("truncate baq_gophers;")

	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()

	Landing(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
	assert.Contains(t, resp.Body.String(), "GophersBAQ")
	assert.Contains(t, resp.Body.String(), "Aún no tenemos ningun Gopher")
}

func TestLandingWithGophers(t *testing.T) {
	DB, _ := db.TestConnection()
	DB.Exec("truncate baq_gophers;")
	defer DB.Exec("truncate baq_gophers;")

	gophers := []models.Gopher{}
	for index := 0; index < 10; index++ {
		gopher := models.Gopher{}
		fako.Fill(&gopher)
		DB.Create(&gopher)
		gophers = append(gophers, gopher)
	}

	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()

	Landing(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
	assert.Contains(t, resp.Body.String(), "GophersBAQ")

	for _, gopher := range gophers {
		assert.Contains(t, resp.Body.String(), gopher.Name)
	}
}
