package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"serv-test/config"
	"serv-test/internal/models"
	"serv-test/internal/validator"
)

type UserForm struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	validator.Validator
}

var userForm = UserForm{}

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

	models.U.FirstName = firstName
	models.U.LastName = lastName
	models.U.Username = username
	models.U.Email = email

	userForm.FirstName = firstName
	userForm.LastName = lastName
	userForm.Username = username
	userForm.Email = email
	userForm.Password = password

	userForm.CheckField(validator.NotBlank(userForm.FirstName), "firstName", "First name can not be blank.")
	userForm.CheckField(validator.NotBlank(userForm.LastName), "lastName", "Last name can not be blank.")
	userForm.CheckField(validator.NotBlank(userForm.Username), "username", "Usename can not be blank.")
	userForm.CheckField(validator.NotBlank(userForm.Email), "email", "Email can not be blank.")
	userForm.CheckField(validator.NotBlank(userForm.Password), "password", "Password can not be blank.")
	userForm.CheckField(validator.CheckUsername(userForm.Username), "username", "Username already exist.")

	models.HashPassword(&password)
	models.U.Password = password

	if userForm.Valid() {
		err = models.StoreCreateUser(&models.U)
		if err != nil {
			ServerError(w, r, err)
			return
		}

		ts, err := template.ParseFiles("web/html/welcome.html", "web/html/t_navbar.html", "web/html/t_logo.html")

		if err != nil {
			ServerError(w, r, err)
			return
		}
		err = ts.ExecuteTemplate(w, "welcome.html", userForm)
		if err != nil {
			ServerError(w, r, err)
		}
		userForm = UserForm{}

	} else {
		ts, err := template.ParseFiles("web/html/createAccount.html", "web/html/t_navbar.html", "web/html/t_logo.html")

		if err != nil {
			ServerError(w, r, err)
			return
		}
		err = ts.ExecuteTemplate(w, "createAccount.html", userForm)
		if err != nil {
			ServerError(w, r, err)
		}
		userForm = UserForm{}

	}

}
