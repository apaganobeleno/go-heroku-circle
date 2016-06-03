package handlers

import (
	"log"
	"net/http"

	"github.com/apaganobeleno/go-heroku-circle/db"
	"github.com/apaganobeleno/go-heroku-circle/models"
)

//Landing is a basic handler for this templo app
func Landing(w http.ResponseWriter, r *http.Request) {
	view, err := views.GetTemplate("welcome.html")
	if err != nil {
		log.Println("Unexpected template err:", err.Error())
		return
	}

	DB, _ := db.Connection()

	gophers := []models.Gopher{}
	DB.Find(&gophers)

	model := struct {
		Gophers []models.Gopher
	}{gophers}

	view.Execute(w, nil, model)
}

func CreateGopher(w http.ResponseWriter, r *http.Request) {
	DB, _ := db.Connection()
	gopher := models.Gopher{}
	gopher.Name = r.FormValue("Name")
	gopher.Company = r.FormValue("Company")
	DB.Create(&gopher)

	http.Redirect(w, r, "/", 302)
}

func DeleteGopher(w http.ResponseWriter, r *http.Request) {
	DB, _ := db.Connection()
	id := r.URL.Query().Get(":id")

	found := models.Gopher{}
	DB.Find(&found).Where("id = ?", id)
	DB.Delete(&found)

	http.Redirect(w, r, "/", 302)
}
