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

	postTitle := r.PostForm.Get("title")
	postContent := r.PostForm.Get("content")

	models.P.Title = postTitle
	models.P.Content = postContent

	err = models.StoreCreatePost(&models.P)
	if err != nil {
		ServerError(w, r, err)
		return
	}

	ts, err := template.ParseFiles("web/html/index.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		ServerError(w, r, err)
	}
	// publish post here afterwards

}
