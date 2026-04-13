package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/index.html")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

}
