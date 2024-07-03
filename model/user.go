package model

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID          int
	Name        string
	Login       string
	Password    string
	UserStorage string
	UserService
}

type UserService interface {
	GetInfoUser() string
	CreateTask()
	RenameTask()
	EditTask()
	DeleteTask()
}

func (u *User) CreateTask() {

}

func (u *User) RenameTask() {

}

func (u *User) EditTask() {

}

func (u *User) DeleteTask() {

}

func (u *User) SetUserStorage(storagePath string) {
	u.UserStorage = storagePath
}

func (u *User) GetInfoUser() string {
	ret := strconv.Itoa(u.ID) + ";" + u.Name + ";" + u.Login + ";" + u.Password + ";" + u.UserStorage

	return ret
}

func constructUser(login string) (*User, error) {
	userInfo, err := os.ReadFile("user/" + login + ".txt")

	if err != nil {
		log.Println("User not found:", err)
		return nil, err
	}

	userInfoStr := string(userInfo)
	userInfoObject := strings.Split(userInfoStr, ";")

	id, _ := strconv.Atoi(userInfoObject[0])
	name := userInfoObject[1]
	password := userInfoObject[3]
	userStorage := userInfoObject[4]

	u := User{ID: id, Name: name, Login: login, Password: password, UserStorage: userStorage}

	return &u, nil
}
