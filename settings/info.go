package settings

import "fmt"

func ViewInfo() {
	info := "\nWelcome to the task maker!"

	fmt.Println(info)
}

func ViewHelp() {
	help := `
Commands:
1) help -- view help list;
2) exit -- exit the program
3) sign_in -- sign in;
4) sign_up -- sign up;
`
	fmt.Println(help)
}
