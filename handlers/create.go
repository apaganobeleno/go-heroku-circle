package handlers

import (
	"net/http"

	"github.com/apaganobeleno/go-heroku-circle/db"
	"github.com/apaganobeleno/go-heroku-circle/models"
)

//CreateGopher ... You guess
func CreateGopher(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	DB, _ := db.Connection()
	gopher := models.Gopher{}
	gopher.Name = r.Form.Get("Name")
	gopher.Company = r.Form.Get("Company")

	DB.Create(&gopher)
	http.Redirect(w, r, "/", 302)
}
