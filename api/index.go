package handlers

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/index.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		ServerError(w, r, err)
	}

}
