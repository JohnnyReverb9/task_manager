package model

type User struct {
	ID          int
	Name        string
	Login       string
	Password    [32]byte
	UserStorage []string
	UserService
}

type UserService interface {
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

func (u *User) DeleteTask() {}
