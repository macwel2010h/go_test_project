package handlers

import (
	"fmt"
	"serv-test/config"
	"serv-test/internal/models"
)

type templateData struct {
	User *models.User
	Post *models.Post
	Feed *models.Posts
}

func PostFeedDisplay() {
	data := templateData{
		User: &models.U,
		Post: &models.P,
		Feed: &models.Ps,
	}

	stmt := ` SELECT * FROM posts `

	postrows, err := config.App.DB.Query(stmt)
	if err != nil {
		return err
	}

	for postrows.Next() {
		var post Post
		if err := postrows.Scan(&post.ID, &post.UserName, &post.Title, &post.Content); err != nil {
			return err
		}
		Ps.Posts = append(Ps.Posts, post)
	}
	fmt.Printf("this is from getpost : %v", Ps.Posts)
}
