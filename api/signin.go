package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"serv-test/config"
	"serv-test/internal/models"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/signIn.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		ServerError(w, r, err)
	}

}

func PostSignInHandler(w http.ResponseWriter, r *http.Request) {
	if config.App.DB == nil {
		http.Error(w, "Database config not configured", http.StatusInternalServerError)
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Print(err)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	models.U, err = models.CheckUserInDatabase(username, password)
	if err != nil {
		ts, err := template.ParseFiles("web/html/wrongLoginRedirect.html")
		if err != nil {
			ServerError(w, r, err)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			ServerError(w, r, err)
		}
		return
	}
	models.P.UserName = username
	ts, err := template.ParseFiles("web/html/home.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.Execute(w, models.U)
	if err != nil {
		ServerError(w, r, err)
	}
}
