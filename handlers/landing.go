package handlers

import (
	"log"
	"net/http"

	"github.com/apaganobeleno/go-heroku-circle/models"
)

//Landing is a basic handler for this templo app
func Landing(w http.ResponseWriter, r *http.Request) {
	view, err := views.GetTemplate("welcome.jet")
	if err != nil {
		log.Println("Unexpected template err:", err.Error())
		return
	}
	model := struct {
		Gophers []models.Gopher
	}{nil}

	view.Execute(w, nil, model)
}
