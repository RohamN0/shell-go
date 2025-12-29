package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		
		string_command := command[:len(command) - 1]
		string_command = strings.Trim(string_command, " ")
		if string_command != "" {
			if string_command == "exit" {
				return
			}

			fmt.Println(string_command + ": command not found")
		}
	} 
}
