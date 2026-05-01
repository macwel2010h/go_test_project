package handlers

import (
	"fmt"
	"serv-test/config"
	"serv-test/internal/models"
)

type templateData struct {
	User        *models.User
	Post        *models.Post
	Feed        *models.Posts
	FieldErrors map[string]string
	Form        any
}

var Data = templateData{
	User: &models.U,
	Post: &models.P,
	Feed: &models.Ps,
}

func PostFeedDisplay() {

	Data.Feed.Posts = nil

	stmt := ` SELECT * FROM posts ORDER BY created_at DESC`

	postrows, err := config.App.DB.Query(stmt)
	if err != nil {
		return
	}

	defer postrows.Close()

	for postrows.Next() {
		var post models.Post
		if err := postrows.Scan(&post.ID, &post.UserName, &post.Title, &post.Content, &post.Created_at); err != nil {
			return
		}
		Data.Feed.Posts = append(Data.Feed.Posts, post)
	}
	fmt.Printf("this is from getpost : %v", Data.Feed.Posts)
}
