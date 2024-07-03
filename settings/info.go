package settings

import "fmt"

func ViewInfo() {
	info := `Welcome to the task maker!`

	fmt.Println(info)
}

func ViewHelp() {
	help := `
Commands:
1) help -- view help list;
2) sign_in -- sign in;
3) sign_up -- sign up;
`
	fmt.Println(help)
}
