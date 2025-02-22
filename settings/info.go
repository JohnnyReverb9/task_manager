package settings

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

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
5) clear -- clear the console;
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
6) edit_task -- edit a task;
7) delete_task -- delete a task;
8) profile -- view profile list;
9) clear -- clear the console;
`
	fmt.Println(help)
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func ClearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Println(err)
	}
}
