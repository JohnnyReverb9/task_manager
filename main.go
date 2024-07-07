package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"task_maker/auth"
	"task_maker/model"
	"task_maker/settings"
)

var u *model.User = nil

func main() {
	var action string

	settings.ViewInfo()
	fmt.Println("Use help command to get more information")

	if !settings.DirExists("./user") {
		err := os.Mkdir("user/", 0777)
		if err != nil {
			fmt.Println("Error creating user/ directory. Create \"user\" folder.")
			log.Fatal(err)
		}
	}

	if !settings.DirExists("./storage") {
		err := os.Mkdir("storage/", 0777)
		if err != nil {
			fmt.Println("Error creating storage/ directory. Create \"storage\" folder.")
			log.Fatal(err)
		}
	}

	if !settings.DirExists("./misc") {
		err := os.Mkdir("misc/", 0777)
		if err != nil {
			log.Fatal("Error creating misc/ directory.")
		}

		idSequencePath := "misc/id_sequence.txt"
		err = os.WriteFile(idSequencePath, []byte("1"), 0666)

		if err != nil {
			log.Fatal("Error creating sequence.")
		}

		allUsersPath := "misc/users.txt"
		err = os.WriteFile(allUsersPath, []byte(""), 0666)

		if err != nil {
			log.Fatal("Error creating users list file.")
		}
	}

	for {
		fmt.Print("Enter command: ")
		_, err := fmt.Scan(&action)

		if err != nil {
			log.Println(err)
		}

		switch action {
		case "help":
			if u == nil {
				settings.ViewHelp()
			} else {
				settings.ViewHelpAuth()
			}
		case "sign_in":
			if u != nil {
				fmt.Println("You are already signed in")
			} else {
				fmt.Println("\nSign in")

				var login, password string

				fmt.Print("Enter login: ")
				_, err := fmt.Scan(&login)
				if err != nil {
					log.Println(err)
				}

				fmt.Print("Enter password: ")
				_, err = fmt.Scan(&password)
				if err != nil {
					log.Println(err)
				}

				login = strings.TrimSpace(login)
				password = strings.TrimSpace(password)
				passwordHash := sha256.Sum256([]byte(password))
				passwordStr := hex.EncodeToString(passwordHash[:])

				u, err = auth.SignIn(login, passwordStr)

				if err != nil {
					log.Println(err)
				} else {
					log.Println("\nWelcome, " + u.Name + "!\n")
				}
			}
		case "sign_up":
			if u != nil {
				fmt.Println("You are already signed up")
			} else {
				fmt.Println("\nSign up:")

				var name, login, password string

				fmt.Print("Enter name: ")
				_, err := fmt.Scan(&name)
				name = strings.TrimSpace(name)
				if err != nil {
					log.Println("Enter a valid name")
				}

				fmt.Print("Enter login: ")
				_, err = fmt.Scan(&login)
				login = strings.TrimSpace(login)
				if err != nil {
					log.Println("Enter a valid login")
				}

				fmt.Print("Enter password: ")
				_, _ = fmt.Scan(&password)
				password = strings.TrimSpace(password)

				var s string

				u, s, err = auth.SignUp(name, login, password)

				log.Println(s + "\n")

				if err != nil {
					log.Println(err)
				}
			}
		case "log_out":
			if u == nil {
				log.Println("You are not logged in!")
			} else {
				u = nil
				log.Println("Logged out successfully")
			}
		case "profile":
			if u == nil {
				log.Println("You are not logged in!")
			} else {
				str := "Profile:\nID: " + strconv.Itoa(u.ID) + "\nName: " + u.Name + "\nLogin: " + u.Login + "\n"
				fmt.Println(str)
			}
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
