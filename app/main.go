package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
	"shell/basic"
	"shell/navigation"
)

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
			case "echo" : basic.Echo(string_command)
			case "type" : basic.Type(string_command)
			case "pwd" :  navigation.PWD(string_command)
			case "cd" :  navigation.CD(string_command)

			default: basic.RunExternalProgram(string_command)
			}
		}
	} 
}
