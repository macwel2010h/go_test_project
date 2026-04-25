package handlers

import (
	"fmt"
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

	// publish post here afterwards

}
