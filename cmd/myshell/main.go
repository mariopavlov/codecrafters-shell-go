package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Name string
	Description string
	Usage string
	Handler func(args []string) error
}

func commandExit(args []string) error {
	if args[1] == "0" {
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
			Handler: commandExit,
		},

	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		userInput := strings.Split(TrimNewLine(rawInput), " ")

		isEcho := strings.HasPrefix(userInput[0], "echo")
		if isEcho {
			echoText := strings.TrimPrefix(userInput[0], "echo ")
			fmt.Println(echoText)
			continue
		}

		if err != nil {
			fmt.Println("Error reading input: " + err.Error())
			os.Exit(1)
		}

		// Get Command
		command, exists := commands[userInput[0]]
		if exists {
			command.Handler(userInput)
		} else {
			fmt.Println(command.Name + ": command not found")
		}
	}

}

func TrimNewLine(prompt string) string {
	return strings.TrimRight(prompt, "\r\n")
}