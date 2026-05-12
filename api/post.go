package handlers

import (
	"fmt"
	"net/http"
	"serv-test/config"
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
	config.App.SessionManager.Put(r.Context(), "flash", "Post created successfully.")

	PostFeedDisplay(w, r)

	http.Redirect(w, r, "/home", 303)
}
