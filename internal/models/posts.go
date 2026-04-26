package models

import (
	"fmt"
	"serv-test/config"
)

type Posts struct {
	ID       int
	UserName string
	Title    string
	Content  string
}

type Pfeed struct {
	Posts_f []Posts
}

var Postfeed = Pfeed{}

var P = Posts{}

func StoreCreatePost(p *Posts) error {

	stmt := ` INSERT INTO posts (username, title, content) VALUES(?,?,?)`

	_, err := config.App.DB.Exec(stmt, p.UserName, p.Title, p.Content)
	if err != nil {
		return err
	}

	return nil
}

func StoreGetPost(p *Posts) error {
	stmt := ` SELECT * FROM posts `

	postrows, err := config.App.DB.Query(stmt)
	if err != nil {
		return err
	}

	for postrows.Next() {
		var post Posts
		if err := postrows.Scan(&post.ID, &post.UserName, &post.Title, &post.Content); err != nil {
			return err
		}
		Postfeed.Posts_f = append(Postfeed.Posts_f, post)
	}
	fmt.Printf("this is from getpost : %v", Postfeed.Posts_f)

	return nil
}
