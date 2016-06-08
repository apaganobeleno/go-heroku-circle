package handlers

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/apaganobeleno/go-heroku-circle/db"
	"github.com/apaganobeleno/go-heroku-circle/models"
	"github.com/bmizerany/pat"
	"github.com/stretchr/testify/assert"
)

func TestDeleteHandler(t *testing.T) {
	DB, _ := db.TestConnection()
	DB.Exec("truncate baq_gophers;")

	gopher := models.Gopher{
		Name:    "Bryan",
		Company: "Wawandco",
	}

	DB.Create(&gopher)

	r, _ := http.NewRequest("POST", "/delete/"+strconv.Itoa(gopher.ID), nil)
	w := httptest.NewRecorder()

	router := pat.New()
	router.Post("/delete/:id", http.HandlerFunc(DeleteGopher))
	router.ServeHTTP(w, r)

	var count int
	DB.Model(&models.Gopher{}).Count(&count)

	assert.Equal(t, w.Code, 302)
	assert.Equal(t, count, 0)
}

func TestDeleteNotExistingHandler(t *testing.T) {
	DB, _ := db.TestConnection()
	DB.Exec("truncate baq_gophers;")

	r, _ := http.NewRequest("POST", "/delete/12", nil)
	w := httptest.NewRecorder()

	router := pat.New()
	router.Post("/delete/:id", http.HandlerFunc(DeleteGopher))
	router.ServeHTTP(w, r)

	var count int
	DB.Model(&models.Gopher{}).Count(&count)
	assert.Equal(t, w.Code, 302)
}
