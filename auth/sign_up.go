package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"task_maker/model"
)

func SignUp(name string, login string, password string) (*model.User, string, error) {
	sequencePath := "misc/id_sequence.txt"
	allUsersPath := "misc/users.txt"
	userPath := "user/"
	storagePath := "storage/"

	_, err := os.ReadFile(userPath + login + ".txt")

	if err == nil {
		// fmt.Println("User already exists")
		return nil, "", fmt.Errorf("user already exists")
	}

	id, err := os.ReadFile(sequencePath)

	if err != nil {
		log.Println("error reading file sequence")
		return nil, "", err
	}

	idInt, _ := strconv.Atoi(string(id))
	passwordHash := sha256.Sum256([]byte(password))
	passwordStr := hex.EncodeToString(passwordHash[:])
	// log.Fatal(passwordHash, passwordStr)

	u := model.User{ID: idInt, Name: name, Login: login, Password: passwordStr}

	idInt += 1
	idStr := strconv.Itoa(idInt)
	err = os.WriteFile(sequencePath, []byte(idStr), 0644)

	if err != nil {
		log.Println("error writing file sequence")
		return nil, "", err
	}

	err = os.WriteFile(userPath+u.Login+".txt", []byte(u.GetInfoUser()), 0666)

	if err != nil {
		log.Println("error writing file user")
		return nil, "", err
	}

	// err = os.WriteFile(storagePath+"/"+u.Login+".txt", []byte(u.GetInfoUser()), 0644)
	err = os.Mkdir(storagePath+u.Login, 0755)

	if err != nil {
		log.Println("error creating storage directory")
		return nil, "", err
	}

	u.SetUserStorage(storagePath + u.Login)

	err = os.WriteFile(allUsersPath, []byte(login), 0644)

	if err != nil {
		log.Println("error adding user to all users list")
	}

	return &u, "user created successfully", nil
}

func validatePassword(password string) error {
	return nil
}
