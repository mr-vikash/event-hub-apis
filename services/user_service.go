package services

import (
	"errors"
	"eventhub/models"
	"eventhub/repository"
	"eventhub/utils"
)

func RegisterUser(name, email, password string) error {
	hashed, err := utils.HashPassword(password)

	if err != nil {
		return err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: hashed,
	}

	return repository.CreateUser(user)
}

func LoginUser(email, password string) (string, error) {
	user, err := repository.GetUserByEmail(email)

	if err != nil {
		return "", errors.New("User not found")
	}

	validPassword := utils.CheckPassword(password, user.Password)

	if !validPassword {
		return "", errors.New("Invalid Password")
	}

	token, err := utils.GenerateJWT(user.Id)

	return token, nil

}
