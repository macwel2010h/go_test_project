package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID         int
	UserName   string
	Title      string
	Content    string
	Created_at time.Time
}

type PostModel struct {
	DB *sql.DB
}

type Posts struct {
	Posts []Post
}

var p = Post{}
var pm = PostModel{}
var Ps = Posts{}

func StoreCreatePost(p *Post) error {

	stmt := ` INSERT INTO posts (username, title, content) VALUES(?,?,?)`

	_, err := pm.DB.Exec(stmt, p.UserName, p.Title, p.Content)
	if err != nil {
		return err
	}

	return nil
}

func StoreGetPost(p *Post) error {

	return nil
}
