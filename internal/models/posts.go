package models

import "serv-test/config"

type Posts struct {
	ID       int
	UserName string
	Title    string
	Content  string
}

var P = Posts{}

func StoreCreatePost(p *Posts) error {

	stmt := ` INSERT INTO posts (username, title, content) VALUES(?,?,?)`

	_, err := config.App.DB.Exec(stmt, p.UserName, p.Title, p.Content)
	if err != nil {
		return err
	}

	return nil
}
