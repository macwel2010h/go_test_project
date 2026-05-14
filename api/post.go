package handlers

import (
	"fmt"
	"net/http"
	"serv-test/config"
	"serv-test/internal/models"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var p = models.Post{}
	var pm = models.PostModel{}

	err := r.ParseForm()
	if err != nil {
		fmt.Print(err)
		return
	}

	p.Title = r.PostForm.Get("title")
	p.Content = r.PostForm.Get("content")

	err = pm.StoreCreatePost(&p)
	if err != nil {
		ServerError(w, r, err)
		return
	}
	config.App.SessionManager.Put(r.Context(), "flash", "Post created successfully.")

	PostFeedDisplay(w, r)

	http.Redirect(w, r, "/home", 303)
}
