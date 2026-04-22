package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"serv-test/internal/models"
)

var userModel *models.UserModel

func SetUserModel(um *models.UserModel) {
	userModel = um
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if userModel == nil {
		http.Error(w, "user model not initialized", http.StatusInternalServerError)
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

	u := models.User{

		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
	}

	u.ID, err = userModel.StoreCreateUser(u)
	if err != nil {
		HTTPError(w, r, err)
		return
	}

	ts, err := template.ParseFiles("web/html/welcome.html")

	if err != nil {
		HTTPError(w, r, err)
		return
	}
	err = ts.Execute(w, u)
	if err != nil {
		HTTPError(w, r, err)
	}
}
