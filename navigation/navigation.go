package navigation

import (
	"fmt"
	"os"
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
