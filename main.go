package main

import (
	"fmt"
	"log"
	"task_maker/auth"
	"task_maker/model"
	"task_maker/settings"
)

var u *model.User

func main() {
	var action string

	settings.ViewInfo()
	fmt.Println("Use help command to get more information")

	for {
		fmt.Print("Enter command: ")
		_, err := fmt.Scan(&action)

		if err != nil {
			log.Println(err)
		}

		switch action {
		case "help":
			settings.ViewHelp()
		case "sign_in":
			// auth.SignIn()
		case "sign_up":
			fmt.Println("\nSign up:")

			var name, login, password string

			fmt.Print("Enter name: ")
			_, err := fmt.Scan(&name)
			if err != nil {
				log.Println("Enter valid name")
			}

			fmt.Print("Enter login: ")
			_, err = fmt.Scan(&login)
			if err != nil {
				log.Println("Enter valid login")
			}

			fmt.Print("Enter password: ")
			_, _ = fmt.Scan(&password)

			var s string

			u, s, err = auth.SignUp(name, login, password)

			log.Println(s + "\n")

			// log.Println(*u)

			if err != nil {
				log.Println(err)
			}
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
