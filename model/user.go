package model

import (
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

func constructUser(login string, userInfo []byte) *User {
	userInfoStr := string(userInfo)
	userInfoObject := strings.Split(userInfoStr, ";")

	id, _ := strconv.Atoi(userInfoObject[0])
	name := userInfoObject[1]
	password := userInfoObject[3]
	userStorage := userInfoObject[4]

	u := User{ID: id, Name: name, Login: login, Password: password, UserStorage: userStorage}

	return &u
}

func ValidateUser(login, password string) (*User, bool) {
	userInfo, err := os.ReadFile("user/" + login + ".txt")

	if err != nil {
		// log.Println("User not found:", err)
		return nil, false
	}

	u := constructUser(login, userInfo)

	if u.Password == password {
		return u, true
	} else {
		return nil, false
	}
}
