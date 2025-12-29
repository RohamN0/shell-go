package main

import (
	"fmt"
	"bufio"
	"os"
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
				fmt.Println(second_command + ": not found")
			}
			
			removed_from_command = strings.Join(removed_from_command_splited[1:], " ")
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

			default: fmt.Println(string_command + ": command not found")
			}
		}
	} 
}
