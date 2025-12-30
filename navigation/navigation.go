package navigation

import (
	"fmt"
	"os"
	"strings"
)

func PWD(command string) {
	// checking if no args has been passed with pwd or not
	if command != "pwd" {
		fmt.Println("pwd: too many arguments")
	} else {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(dir)
		}
	}
}

func CD(command string) {
	splited_command := strings.Fields(command)
	if len(splited_command) == 1 {
		os.Chdir("/")
	} else if len(splited_command) > 2 {
		fmt.Println("cd: too many arguments")
	} else {
		path := splited_command[1]
		_, err := os.Stat(path)

		if err == nil {
			os.Chdir(path)
		} else if strings.Contains(path, "~") {
			users_home_dir, _ := os.UserHomeDir()
			abs_path := strings.ReplaceAll(path, "~", users_home_dir)
			_, err := os.Stat(abs_path)
			
			if os.IsNotExist(err) {
				fmt.Println("cd:", path, ": No such file or directory")
			} else {
				os.Chdir(abs_path)
			}

		} else if os.IsNotExist(err) {
			fmt.Println("cd:", path, ": No such file or directory")
		}
	}
}
