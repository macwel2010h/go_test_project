package handlers

import (
	"html/template"
	"net/http"
)

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/createAccount.html")
	if err != nil {
		HTTPError(w, r, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		HTTPError(w, r, err)
	}

}
