package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"serv-test/internal/models"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Print(err)
		return
	}

	models.P.Title = r.PostForm.Get("title")
	models.P.Content = r.PostForm.Get("content")

	err = models.StoreCreatePost(&models.P)
	if err != nil {
		ServerError(w, r, err)
		return
	}

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
