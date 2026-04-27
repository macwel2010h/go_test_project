package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"serv-test/config"
	"serv-test/internal/models"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/signIn.html", "web/html/t_navbar.html", "web/html/t_logo.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "signIn.html", Data)
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
		ts, err := template.ParseFiles("web/html/wrongLoginRedirect.html", "web/html/t_navbar.html", "web/html/t_logo.html")
		if err != nil {
			ServerError(w, r, err)
			return
		}
		err = ts.ExecuteTemplate(w, "wrongLoginRedirection.html", nil)
		if err != nil {
			ServerError(w, r, err)
		}
		return
	}

	models.P.UserName = username
	PostFeedDisplay()

	ts, err := template.ParseFiles("web/html/home.html", "web/html/t_navbar.html", "web/html/t_logo.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "home.html", Data)
	if err != nil {
		ServerError(w, r, err)
	}
}
