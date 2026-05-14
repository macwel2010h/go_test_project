package handlers

import (
	"html/template"
	"net/http"
	"serv-test/config"
	"serv-test/helpers"
	"serv-test/internal/models"
	"serv-test/internal/validator"
)

type UserForm struct {
	FirstName           string `form:"firstName"`
	LastName            string `form:"lastName"`
	Username            string `form:"username"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

var userForm = UserForm{}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if config.App.DB == nil {
		http.Error(w, "Database config not configured", http.StatusInternalServerError)
		return
	}
	err := helpers.DecodeForm(r, &userForm)
	if err != nil {
		ClientError(w, http.StatusBadRequest)
		return
	}

	var u = models.User{}
	var um = models.UserModel{}

	u.FirstName = r.PostForm.Get("firstName")
	u.LastName = r.PostForm.Get("lastName")
	u.Username = r.PostForm.Get("username")
	u.Email = r.PostForm.Get("email")
	u.Password = r.PostForm.Get("password")

	models.HashPassword(&u.Password)

	userForm.CheckField(validator.NotBlank(userForm.FirstName), "firstName", "First name can not be blank.")
	userForm.CheckField(validator.NotBlank(userForm.LastName), "lastName", "Last name can not be blank.")
	userForm.CheckField(validator.NotBlank(userForm.Username), "username", "Usename can not be blank.")
	userForm.CheckField(validator.NotBlank(userForm.Email), "email", "Email can not be blank.")
	userForm.CheckField(validator.NotBlank(userForm.Password), "password", "Password can not be blank.")
	userForm.CheckField(validator.CheckUsername(userForm.Username), "username", "Username already exist.")

	if userForm.Valid() {
		err = um.StoreCreateUser(u)
		if err != nil {
			ServerError(w, r, err)
			return
		}

		http.Redirect(w, r, "/welcome", 303)

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
