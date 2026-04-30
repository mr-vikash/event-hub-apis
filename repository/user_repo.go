package repository

import (
	"eventhub/config"
	"eventhub/models"
)

func CreateUser(user models.User) error {
	query := "insert into users(name,email,password) values(?,?,?)"
	_, err := config.DB.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	row := config.DB.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
