package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"serv-test/config"
	"serv-test/internal/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if config.App.DB == nil {
		http.Error(w, "Database config not configured", http.StatusInternalServerError)
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Print(err)
		return
	}

	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")
	username := r.PostForm.Get("username")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	models.HashPassword(&password)

	models.U.FirstName = firstName
	models.U.LastName = lastName
	models.U.Username = username
	models.U.Email = email
	models.U.Password = password

	err = models.StoreCreateUser(&models.U)
	if err != nil {
		ServerError(w, r, err)
		return
	}

	ts, err := template.ParseFiles("web/html/welcome.html")

	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.Execute(w, models.U)
	if err != nil {
		ServerError(w, r, err)
	}
}
