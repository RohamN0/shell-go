package basic

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"shell/tokenizer"
)

func Echo(command string) {
	command_splited := tokenizer.Tokenize(command)[1:]
	for i, text := range(command_splited) {
		command_splited[i] = tokenizer.EchoTokenize(text)
	}
	command = strings.Join(command_splited, " ")

	fmt.Println(command)
}

func Type(command string) {
	commands := tokenizer.Tokenize(command)[1:]
	valid_commands := [] string{"echo", "exit", "type", "pwd", "cd"}

	for _, command := range(commands) {
		flag := false
		for _, valid_command := range(valid_commands) {
			// cheking of the text is a valid command or not
			if valid_command == command {
				fmt.Println(command + " is a shell builtin")
				flag = true
				break
			}
		}

		if !flag {
			path, err := exec.LookPath(command)
			if err != nil {
				fmt.Println(command + ": not found")
			} else {
				fmt.Println(command + " is " + path)
			}
		}
	}
}

func RunExternalProgram(command string) {
    args := tokenizer.Tokenize(command)
    prog := args[0]

	_, err := exec.LookPath(prog)
    if err != nil {
		fmt.Println(prog + ": command not found")
    } else {
		cmd := exec.Command(prog, args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.Run()
	}
}
