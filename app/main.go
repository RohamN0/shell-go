package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func echo(command string) {
	_, text, _ := strings.Cut(command, "echo ")
	fmt.Println(text)
}

func type_command(command string) {
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

func run_external_program(path string) {
	// checking if there is only one command with no args
	command, err := exec.LookPath(path)
	if err == nil {
		cmd := exec.Command(command)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
    		fmt.Fprintf(os.Stderr, "%v\n", err)
		}

	} else {
		command := strings.Split(path, " ")[0]

		actual_command, err := exec.LookPath(command)
		if err != nil {
			fmt.Println(command + ": command not found")
		} else {
			args := [] string{actual_command}
			removed_from_command := strings.Join(strings.Split(path, " ")[1:], " ")

			for {
				// using trim and split to remove all the spaces and get the actual text each iteration
				removed_from_command = strings.Trim(removed_from_command, " ")
				removed_from_command_splited := strings.Split(removed_from_command, " ")
				arg := removed_from_command_splited[0]
				fmt.Println(arg)
				// appending all the args
				args = append(args, arg)
				
				// break statement
				if arg == "" {
					break
				}
				
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

func main() {
	for {
		fmt.Print("$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		
		string_command := command[:len(command) - 1]
		string_command = strings.Trim(string_command, " ")
		if string_command != "" {
			first_command := strings.Split(string_command, " ")[0]
			switch first_command {

			case "exit" : return
			case "echo" : echo(string_command)
			case "type" : type_command(string_command)

			default: run_external_program(string_command)
			}
		}
	} 
}
