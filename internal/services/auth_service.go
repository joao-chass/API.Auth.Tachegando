package services

import (
	"AuthTachegando/internal/models"
	"AuthTachegando/internal/repositories"
	"AuthTachegando/pkg/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{Username: username, Password: string(hashedPassword)}
	return repositories.CreateUser(user)
}

func AuthenticateUser(username, password string) (string, error) {
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("usuário não encontrado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("senha inválida")
	}

	token, err := jwt.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
