package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/util"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {

	query := `INSERT INTO USERS(email, password) VALUES(?, ?)`

	preparedQuery, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	hashedPassword, err := util.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := preparedQuery.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email=?"
	row := db.DB.QueryRow(query, u.Email)

	var retreivePassword string
	err := row.Scan(&u.ID, &retreivePassword)

	if err != nil {
		return errors.New("INVALID CREDENTIALS")
	}

	passwordIsValid := util.CheckPasswordHash(u.Password, retreivePassword)

	if !passwordIsValid {
		return errors.New("INVALID CREDENTIALS")
	}
	return nil
}
