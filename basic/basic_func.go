package basic

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Echo(command string) {
	_, text, _ := strings.Cut(command, "echo ")
	fmt.Println(text)
}

func Type(command string) {
	if command == "type" {
		fmt.Println("type is a shell builtin")
	} else {
		removed_from_command := strings.Join(strings.Split(command, "type ")[1:], " ")
		valid_commands := [] string{"echo", "exit", "type"}

		for {
			// using trim and split to remove all the spaces and get the actual text each iteration
			removed_from_command = strings.Trim(removed_from_command, " ")
			removed_from_command_splited := strings.Split(removed_from_command, " ")
			second_command := removed_from_command_splited[0]

			// break statement
			if second_command == "" {
				break
			}
			
			flag := false
			for _, valid_command := range(valid_commands) {
				// cheking of the text is a valid command or not
				if valid_command == second_command {
					fmt.Println(second_command + " is a shell builtin")
					flag = true
					break
				}
			}

			if !flag {
				path, err := exec.LookPath(second_command)
				if err != nil {
					fmt.Println(second_command + ": not found")
				} else {
					fmt.Println(second_command + " is " + path)
				}
			}
			
			removed_from_command = strings.Join(removed_from_command_splited[1:], " ")
		}
	}
}

func RunExternalProgram(command string) {
	// checking if there is only one command with no args
	_, err := exec.LookPath(command)
	if err == nil {
		cmd := exec.Command(command)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
    		fmt.Fprintf(os.Stderr, "%v\n", err)
		}

	} else {
		actual_command := strings.Split(command, " ")[0]

		_, err := exec.LookPath(actual_command)
		if err != nil {
			fmt.Println(actual_command + ": command not found")
		} else {
			args := [] string{actual_command}
			removed_from_command := strings.Join(strings.Split(command, " ")[1:], " ")

			for {
				// using trim and split to remove all the spaces and get the actual text each iteration
				removed_from_command = strings.Trim(removed_from_command, " ")
				removed_from_command_splited := strings.Split(removed_from_command, " ")
				arg := removed_from_command_splited[0]
				
				// break statement
				if arg == "" {
					break
				}
				// appending all the args
				args = append(args, arg)
				
				removed_from_command = strings.Join(removed_from_command_splited[1:], " ")
			}

			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
    			fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		}
	}
}
