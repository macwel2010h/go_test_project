package handlers

import (
	"html/template"
	"net/http"
)

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/createAccount.html", "web/html/t_navbar.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "createAccount.html", nil)
	if err != nil {
		ServerError(w, r, err)
	}

}
