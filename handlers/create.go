package handlers

import "net/http"

func CreateGopher(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.WriteHeader(201)

}
