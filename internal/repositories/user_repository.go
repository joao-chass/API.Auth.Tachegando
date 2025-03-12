package repositories

import (
	"AuthTachegando/internal/db"
	"AuthTachegando/internal/models"
	"log"
)

func CreateUser(user models.User) error {
	_, err := db.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		log.Println("Erro ao criar usuário:", err)
		return err
	}
	return nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.DB.QueryRow("SELECT id, username, password FROM users WHERE username=$1", username).
		Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
