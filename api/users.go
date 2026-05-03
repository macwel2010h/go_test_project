package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"serv-test/config"
	"serv-test/internal/models"
	"strings"
)

type UserForm struct {
	FirstName   string
	LastName    string
	Username    string
	Email       string
	Password    string
	FieldErrors map[string]string
}

var userForm = UserForm{

	FieldErrors: map[string]string{},
}

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

	userForm.FirstName = firstName
	userForm.LastName = lastName
	userForm.Username = username
	userForm.Email = email
	userForm.Password = password

	if strings.TrimSpace(firstName) == "" {
		userForm.FieldErrors["firstName"] = "First name can not be empty."
	}

	if strings.TrimSpace(lastName) == "" {
		userForm.FieldErrors["lastName"] = "Last name can not be empty."
	}

	if strings.TrimSpace(username) == "" {
		userForm.FieldErrors["username"] = "Username can not be empty."
	} else if models.CheckUsernameAvailability(username) == true {
		userForm.FieldErrors["username"] = "Username already exist. select different username."

	}

	if strings.TrimSpace(email) == "" {
		userForm.FieldErrors["email"] = "Email can not be empty."
	}

	if strings.TrimSpace(password) == "" {
		userForm.FieldErrors["password"] = "Password can not be empty."
	}

	if len(userForm.FieldErrors) > 0 {
		ts, err := template.ParseFiles("web/html/createAccount.html", "web/html/t_navbar.html", "web/html/t_logo.html")

		if err != nil {
			ServerError(w, r, err)
			return
		}
		err = ts.ExecuteTemplate(w, "createAccount.html", userForm)
		if err != nil {
			ServerError(w, r, err)
		}
	} else {

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
	}
}
