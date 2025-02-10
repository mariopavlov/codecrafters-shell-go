package main

import (
	"fmt"
	"strings"

	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands/concrete"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Handler     func(args []string)
}

var allCommands = make(map[string]Command)

func registerCommands() {
	allCommands["exit"] = Command{
		Name:        "exit",
		Description: "exit is a shell builtin",
		Usage:       "exit [code]",
		Handler:     handleExit,
	}
	allCommands["echo"] = Command{
		Name:        "echo",
		Description: "echo is a shell builtin",
		Usage:       "echo [text to print out]",
		Handler:     handleEcho,
	}
	allCommands["type"] = Command{
		Name:        "type",
		Description: "type is a shell builtin",
		Usage:       "type [command]",
		Handler:     handleType,
	}
	allCommands["pwd"] = Command{
		Name:        "pwd",
		Description: "pwd is a shell builtin",
		Usage:       "pwd [command]",
		Handler:     handlePwd,
	}
	allCommands["cd"] = Command{
		Name:        "cd",
		Description: "Change current working directory",
		Usage:       "cd [directory]",
		Handler:     handleChangeDirectory,
	}
}

func handleChangeDirectory(args []string) {
	if len(args) <= 1 {
		return
	}

	directory := args[1]

	command := commands.NewChangeDirectoryCommand(directory)
	command.Execute()
}

func handlePwd(args []string) {
	command := commands.NewPwdCommand()
	command.Execute()
}

func handleExit(args []string) {
	command := commands.NewExitCommand()
	command.Execute()
}

func handleEcho(args []string) {
	var message string = strings.Join(args[1:], " ")

	command := commands.NewEchoCommand(message)
	command.Execute()

}

func handleType(args []string) {
	if len(args) <= 1 {
		fmt.Println("Error: Type command requires command as a parameter")
	}

	commandToDescribe := args[1]

	command := commands.NewTypeCommand(commandToDescribe)

	internalCommand, exists := allCommands[commandToDescribe]

	if exists {
		fmt.Println(internalCommand.Description)
	} else {
		command.Execute()
	}
}

func GetCommand(command string) (Command, bool) {
	value, isExist := allCommands[command]

	return value, isExist
}

func init() {
	registerCommands()
}
