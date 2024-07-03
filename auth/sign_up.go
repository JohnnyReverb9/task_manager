package auth

import (
	"crypto/sha256"
	"log"
	"os"
	"strconv"
	"task_maker/model"
)

func SignUp(name string, login string, password string) (*model.User, string, error) {
	id, err := os.ReadFile("misc/id_sequence.txt")

	if err != nil {
		log.Println("Error reading file sequence")
	}

	idInt, _ := strconv.Atoi(string(id))
	passwordHash := sha256.Sum256([]byte(password))

	u := model.User{ID: idInt, Name: name, Login: login, Password: passwordHash}

	return &u, "User created successfully", nil
}

func validatePassword(password string) error {
	return nil
}
