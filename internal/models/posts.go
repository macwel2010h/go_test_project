package models

import (
	"serv-test/config"
)

type Post struct {
	ID         int
	UserName   string
	Title      string
	Content    string
	Created_at []uint8
}

type Posts struct {
	Posts []Post
}

var Ps = Posts{}

var P = Post{}

func StoreCreatePost(p *Post) error {

	stmt := ` INSERT INTO posts (username, title, content) VALUES(?,?,?)`

	_, err := config.App.DB.Exec(stmt, p.UserName, p.Title, p.Content)
	if err != nil {
		return err
	}

	return nil
}

func StoreGetPost(p *Post) error {

	return nil
}
