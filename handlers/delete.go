package handlers

import (
	"net/http"

	"github.com/apaganobeleno/go-heroku-circle/db"
	"github.com/apaganobeleno/go-heroku-circle/models"
)

func DeleteGopher(w http.ResponseWriter, r *http.Request) {
	DB, _ := db.Connection()
	id := r.URL.Query().Get(":id")

	DB.Delete(&models.Gopher{}, "id = ?", id)
	http.Redirect(w, r, "/", 302)
}
