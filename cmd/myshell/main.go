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
	if len(args) <= 1 {
		fmt.Println("Error: Exit requires exit code ")
		os.Exit(1)
	}

	if args[1] == "0" {
		os.Exit(0)
	} else {
		fmt.Println("Error: Exit with code " + args[1])
		os.Exit(1)
	}

	return nil
}

func commandEcho(args []string) error {

	echoText := strings.Join(args[1:], " ")
	fmt.Println(args)
	fmt.Println(echoText)

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
		"echo": {
			Name: "echo",
			Description: "echo is a shell builtin",
			Usage: "echo [text to print out]",
			Handler: commandEcho,
		},
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		userInput := strings.Split(TrimNewLine(rawInput), " ")

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