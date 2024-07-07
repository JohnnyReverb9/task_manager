package main

import (
	"bufio"
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

	// START INFO
	settings.ViewInfo()
	fmt.Println("Use help command to get more information")

	// DIR CHECKS
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

	// COMMAND CYCLE
	for {
		fmt.Print("Enter command: ")
		_, err := fmt.Scan(&action)

		if err != nil {
			log.Println(err)
		}

		switch action {
		// VIEW HELP
		case "help":
			if u == nil {
				settings.ViewHelp()
			} else {
				settings.ViewHelpAuth()
			}
		// AUTH BLOCK
		// SIGN IN BLOCK
		case "sign_in":
			if u != nil {
				fmt.Println("You are already signed in")
				break
			}
			fmt.Println("\nSign in:")
			reader := bufio.NewReader(os.Stdin)

			var login, password string

			fmt.Print("Enter login: ")
			login, err = reader.ReadString('\n')
			if err != nil {
				log.Println(err)
			}

			fmt.Print("Enter password: ")
			password, err = reader.ReadString('\n')
			if err != nil {
				log.Println(err)
			}

			login = strings.TrimSpace(login)
			password = strings.TrimSpace(password)
			passwordHash := sha256.Sum256([]byte(password))
			passwordStr := hex.EncodeToString(passwordHash[:])

			u, err = auth.SignIn(login, passwordStr)

			settings.ClearConsole()

			if err != nil {
				log.Println(err)
			} else {
				log.Println("\nWelcome, " + u.Name + "!\n")
			}
		// SIGN UP BLOCK
		case "sign_up":
			if u != nil {
				fmt.Println("You are already signed up")
				break
			}
			fmt.Println("\nSign up:")

			var name, login, password string
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter name: ")
			name, err = reader.ReadString('\n')

			name = strings.TrimSpace(name)
			if err != nil {
				log.Println("Enter a valid name")
			}

			fmt.Print("Enter login: ")
			login, err = reader.ReadString('\n')
			login = strings.TrimSpace(login)
			if err != nil {
				log.Println("Enter a valid login")
			}

			fmt.Print("Enter password: ")
			password, err = reader.ReadString('\n')
			password = strings.TrimSpace(password)

			var s string

			u, s, err = auth.SignUp(name, login, password)

			settings.ClearConsole()

			log.Println(s + "\n")

			if err != nil {
				log.Println(err)
			}
		// LOG OUT
		case "log_out":
			if u == nil {
				log.Println("You are not logged in!")
			} else {
				u = nil
				log.Println("Logged out successfully")
			}
		// END AUTH BLOCK

		// TASK MANAGEMENT BLOCK
		// CREATE TASK
		case "create_task":
			if u == nil {
				log.Println("You are not logged in!")
				break
			}

			var taskName, content string
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\nCreate task:")

			fmt.Print("Enter task name: ")
			taskName, err = reader.ReadString('\n')
			taskName = strings.TrimSpace(taskName)
			if err != nil {
				log.Println("Enter a valid task name")
			}

			fmt.Print("Enter description: ")
			content, err = reader.ReadString('\n')
			// content = strings.TrimSpace(content)
			if err != nil {
				log.Println("Enter a valid task content")
			}

			u.CreateTask(taskName, content)

			log.Println("Task \"" + taskName + "\" created successfully")
		// EDIT TASK
		case "edit_task":
			if u == nil {
				log.Println("You are not logged in!")
				break
			}
		// DELETE TASK
		case "delete_task":
			if u == nil {
				log.Println("You are not logged in!")
				break
			}
		// VIEW ALL TASKS
		case "view_tasks":
			if u == nil {
				log.Println("You are not logged in!")
				break
			}
			u.ViewTasks()
		// END TASK MANAGEMENT BLOCK
		// MISC
		case "profile":
			if u == nil {
				log.Println("You are not logged in!")
			} else {
				str := "Profile:\nID: " + strconv.Itoa(u.ID) + "\nName: " + u.Name + "\nLogin: " + u.Login + "\n"
				fmt.Println(str)
			}
		case "clear":
			settings.ClearConsole()
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
