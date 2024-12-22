package models

import (
	"errors"
	"github.com/npinnaka/goproject/db"
	"github.com/npinnaka/goproject/utils"
	"log"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Save() error {
	var err error
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	insertStatement := `INSERT INTO users (email,password) VALUES (?, ?)`
	result, err := db.DB.Exec(insertStatement, u.Email, u.Password)
	if err != nil {
		return err
	}
	u.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update() (*User, error) {
	password, err := utils.HashPassword(u.Password)
	updateStatement := `UPDATE users SET password = ?, email = ? WHERE id = ?`
	results, err := db.DB.Exec(updateStatement, password, u.Email, u.ID)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("event not found")
	}
	return u, nil
}

func (u *User) FindUser() (*string, error) {
	var dbUserPassword string
	row := db.DB.QueryRow("SELECT password FROM users where email = ?", u.Email)
	err := row.Scan(&dbUserPassword)
	if err != nil {
		return nil, err
	}
	if !utils.CompareHashPassword(dbUserPassword, u.Password) {
		return nil, nil
	}
	signedToken, err := utils.GenerateJWTToken(u.Email, u.ID)
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}

func GetAllUsers() ([]User, error) {
	rows, err := db.DB.Query("SELECT id, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserById(id int64) (*User, error) {
	row := db.DB.QueryRow("SELECT id, email FROM users where id = ?", id)
	var event User

	err := row.Scan(&event.ID, &event.Email)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func DeleteUserById(id int64) (*int64, error) {
	deleteUserQuery := `DELETE FROM users WHERE id = ?`
	result, err := db.DB.Exec(deleteUserQuery, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if affected == 0 {
		return nil, errors.New("User not found")
	}
	return &affected, nil
}
