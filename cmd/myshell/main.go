package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	commands := make(map[string]string)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')

		// Trim newline
		userInput = userInput[:len(userInput)-1]

		if userInput == "exit 0" {
			os.Exit(0)
		}

		if err != nil {
			fmt.Println("Error reading input: " + err.Error())
			os.Exit(1)
		}

		// Get Command
		command, exists := commands[userInput]
		if !exists {
			command = userInput
		}

		if exists {
			fmt.Println(command + ": to be executed")
		} else {
			fmt.Println(command + ": command not found")
		}
	}

}
