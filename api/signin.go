package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/signIn.html")
	if err != nil {
		HTTPError(w, r, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		HTTPError(w, r, err)
	}

}

func PostSignInHandler(w http.ResponseWriter, r *http.Request) {
	if userModel == nil {
		http.Error(w, "user model not initialized", http.StatusInternalServerError)
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Print(err)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	id, err := userModel.CheckUserInDatabase(username, password)
	if err != nil {
		HTTPError(w, r, err)
		return
	}

	if id == 0 {

		ts, err := template.ParseFiles("web/html/wrongLoginRedirect.html")
		if err != nil {
			HTTPError(w, r, err)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			HTTPError(w, r, err)
		}
	} else {
		ts, err := template.ParseFiles("web/html/home.html")
		if err != nil {
			HTTPError(w, r, err)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			HTTPError(w, r, err)
		}
	}

}
