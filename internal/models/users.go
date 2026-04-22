package models

import (
	"database/sql"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	CreatedAt []uint8
}

type UserModel struct {
	DB *sql.DB
}

var U = User{}

func (um *UserModel) StoreCreateUser(u *User) (int, error) {
	stmt := `INSERT INTO users (firstName, lastName, username, email, password)
	VALUES(?,?,?,?,?)`

	result, err := um.DB.Exec(stmt, u.FirstName, u.LastName, u.Username, u.Email, u.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (um *UserModel) StoreDeleteUser(username string) (int, error) {
	return 0, nil
}

func (um *UserModel) CheckUserInDatabase(username, password string) (User, error) {
	stmt := `SELECT * FROM users WHERE username = ? AND password = ?`

	err := um.DB.QueryRow(stmt, username, password).Scan(&U.ID, &U.FirstName, &U.LastName, &U.Username, &U.Email, &U.Password, &U.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil
		}
		return User{}, err
	}

	return U, nil
}
