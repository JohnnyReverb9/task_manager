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
2) exit -- exit the program;
3) sign_in -- sign in;
4) sign_up -- sign up;
`
	fmt.Println(help)
}

func ViewHelpAuth() {
	help := `
Commands:
1) help -- view help list;
2) exit -- exit the program;
3) log_out -- log out;
4) view_tasks -- view all tasks;
5) create_task -- create a new task;
6) rename_task -- rename a task;
7) delete_task -- delete a task;
8) profile -- view profile list;
`
	fmt.Println(help)
}
