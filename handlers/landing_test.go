package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/apaganobeleno/go-heroku-circle/db"
	"github.com/apaganobeleno/go-heroku-circle/models"
	"github.com/bmizerany/pat"
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

func TestCreateGopher(t *testing.T) {
	req, _ := http.NewRequest("POST", "/create", nil)
	req.Form = url.Values{}
	req.Form.Set("Name", "SQL")
	req.Form.Set("Company", "Wawandco")

	DB, _ := db.TestConnection()
	var count int
	DB.Model(&models.Gopher{}).Count(&count)

	rw := httptest.NewRecorder()
	CreateGopher(rw, req)

	var countAfter int
	DB.Model(&models.Gopher{}).Count(&countAfter)

	assert.Equal(t, rw.Code, 302)
	assert.Equal(t, countAfter, count+1)
}

func TestDeleteGopher(t *testing.T) {
	DB, _ := db.TestConnection()

	gopher := models.Gopher{}
	fako.Fill(&gopher)
	DB.Create(&gopher)

	id := gopher.ID
	route := "/delete/" + strconv.Itoa(id)

	req, _ := http.NewRequest("POST", route, nil)
	rw := httptest.NewRecorder()

	router := pat.New()
	router.Post("/delete/:id", http.HandlerFunc(DeleteGopher))
	router.ServeHTTP(rw, req)

	var found models.Gopher
	DB.Model(&models.Gopher{}).Find(&gopher).Where("id = ?", id)

	assert.Equal(t, rw.Code, 302)
	assert.Equal(t, found.ID, 0)

}

func TestDeleteGopherNotExisting(t *testing.T) {
	req, _ := http.NewRequest("POST", "/delete/200", nil)
	rw := httptest.NewRecorder()

	router := pat.New()
	router.Post("/delete/:id", http.HandlerFunc(DeleteGopher))
	router.ServeHTTP(rw, req)

	assert.Equal(t, rw.Code, 302)
}
