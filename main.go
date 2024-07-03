package main

import (
	"fmt"
	"log"
	"task_maker/settings"
)

func main() {
	var action string

	settings.ViewInfo()
	fmt.Println("Use help command to get more information")

	for {
		fmt.Println("Enter command: ")
		_, err := fmt.Scan(&action)

		if err != nil {
			log.Println(err)
		}

		switch action {
		case "help":
			settings.ViewHelp()
		default:
			fmt.Println("Unknown command")
		}
	}
}
