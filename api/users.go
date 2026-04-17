package handlers

import (
	"fmt"
	"net/http"
	"serv-test/internal/models"
	"strconv"
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

	id, err := userModel.StoreCreateUser(firstName, lastName, username, email, password)
	if err != nil {
		HTTPError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("created user id: " + strconv.Itoa(id)))
}
