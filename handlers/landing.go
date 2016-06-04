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

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 302)
}
