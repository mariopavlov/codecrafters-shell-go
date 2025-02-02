package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

type Command struct {
	Name string
	Description string
	Usage string
	Handler func(args []string) error
}

func exit(args []string) error {
	if args[0] == "0" {
		os.Exit(0)
	} else {
		fmt.Println("Error: Exit with code " + args[1])
		os.Exit(1)
	}

	return nil
}

func main() {
	commands := map[string] Command {
		"exit": {
			Name: "exit",
			Description: "exit is a shell builtin",
			Usage: "exit [code]",
			Handler: exit,
		},

	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		userInput := TrimNewLine(rawInput)

		isEcho := strings.HasPrefix(userInput, "echo")
		if isEcho {
			echoText := strings.TrimPrefix(userInput, "echo ")
			fmt.Println(echoText)
			continue
		}

		if err != nil {
			fmt.Println("Error reading input: " + err.Error())
			os.Exit(1)
		}

		// Get Command
		command, exists := commands[userInput]
		if !exists {
			fmt.Println(command.Name + ": command not found")
		}

		if exists {
			fmt.Println(command.Name + ": to be executed")
		} else {
			
		}
	}

}

func TrimNewLine(prompt string) string {
	return strings.TrimRight(prompt, "\r\n")
}