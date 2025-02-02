package main

import (
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

var commands = make(map[string]Command)

func registerCommands() {
    commands["exit"] = Command{
        Name:        "exit",
        Description: "exit is a shell builtin",
        Usage:       "exit [code]",
        Handler:     commandExit,
    }
    commands["echo"] = Command{
        Name:        "echo",
        Description: "echo is a shell builtin",
        Usage:       "echo [text to print out]",
        Handler:     commandEcho,
    }
    commands["type"] = Command{
        Name:        "type",
        Description: "type is a shell builtin",
        Usage:       "type [command]",
        Handler:     commandType,
    }
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
	fmt.Println(echoText)

	return nil
}

func commandType(args []string) error {
	if len(args) <= 1 {
		fmt.Println("Error: Type command requires command as a parameter")
		return nil
	}

	describeCommand := args[1]

	command, exists := commands[describeCommand]

	if exists {
		fmt.Println(command.Description)

	} else {
		fmt.Println(describeCommand + ": not found")
	}
	return nil
}

func GetCommand(command string) (Command, bool) {
	value, isExist := commands[command]

	return value, isExist
}

func init() {
	registerCommands()
}