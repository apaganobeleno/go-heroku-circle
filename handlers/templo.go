package handlers

import (
	"log"
	"net/http"
)

//Templo is a sample handler to show off some of the components of the template
func Templo(w http.ResponseWriter, r *http.Request) {
	view, err := views.GetTemplate("templo.jet")
	if err != nil {
		log.Println("Unexpected template err:", err.Error())
		return
	}

	name := r.URL.Query().Get(":name")
	view.Execute(w, nil, struct {
		Name string
	}{name})
}
