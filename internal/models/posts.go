package models

import (
	"serv-test/config"
)

type Posts struct {
	ID       int
	UserName string
	Title    string
	Content  string
}

type Pfeed struct {
	Posts []Posts
}

var pfeed = Pfeed{}

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

	postrows.Scan(pfeed)

	return nil
}
