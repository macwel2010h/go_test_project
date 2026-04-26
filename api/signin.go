package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"serv-test/config"
	"serv-test/internal/models"
)

type templateData struct {
	User *models.User
	Post *models.Posts
	Feed *models.Pfeed
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/signIn.html", "web/html/t_navbar.html", "web/html/t_logo.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "signIn.html", nil)
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
	models.StoreGetPost(&models.P)
	models.P.UserName = username

	data := templateData{
		User: &models.U,
		Post: &models.P,
		Feed: &models.Postfeed,
	}

	fmt.Println(data.Feed.Posts_f)

	ts, err := template.ParseFiles("web/html/home.html", "web/html/t_navbar.html", "web/html/t_logo.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		ServerError(w, r, err)
	}
}
