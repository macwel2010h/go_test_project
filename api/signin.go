package handlers

import (
	"html/template"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/signIn.html")
	if err != nil {
		HTTPError(w, r, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		HTTPError(w, r, err)
	}

}
