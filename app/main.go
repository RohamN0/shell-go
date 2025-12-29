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
		if strings.Trim(string_command, " ")  != "" {
			fmt.Println(string_command + ": command not found")
		}
	} 
}
