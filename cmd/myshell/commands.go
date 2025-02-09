package main

import (
	"fmt"
	"os"
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
		Handler:     commandExit,
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
		Handler:     commandType,
	}
	allCommands["pwd"] = Command{
		Name:        "pwd",
		Description: "pwd is a shell builtin",
		Usage:       "pwd [command]",
		Handler:     commandPwd,
	}
	allCommands["cd"] = Command{
		Name:        "cd",
		Description: "Change current working directory",
		Usage:       "cd [directory]",
		Handler:     commandDirectoryChange,
	}
}

func commandDirectoryChange(args []string) {
	if len(args) <= 1 {
		return
	}

	if args[1] == "~" {
		// TODO Do we have case to fail? What is the case that can fail here?
		homeDirectory, _ := os.UserHomeDir()
		os.Chdir(homeDirectory)
	}

	error := os.Chdir(args[1])

	if error != nil {
		fmt.Printf("cd: %v: No such file or directory\n", args[1])
	}

}

// TODO This method does not need args, but need to satisfy the command handler interface
// improve this code
func commandPwd(args []string) {
	currentDir, error := os.Getwd()

	if error == nil {
		fmt.Println(currentDir)
	}

}

func commandExit(args []string) {
	if len(args) <= 1 {
		// Use Default Exit of exit code is missing
		os.Exit(0)
	}

	if args[1] == "0" {
		os.Exit(0)
	} else {
		fmt.Println("Error: Exit with code " + args[1])
		os.Exit(1)
	}
}

func handleEcho(args []string) {
	var message string = strings.Join(args[1:], " ")

	command := commands.NewEchoCommand(message)
	command.Execute()

}

func commandType(args []string) {
	if len(args) <= 1 {
		fmt.Println("Error: Type command requires command as a parameter")
	}

	describeCommand := args[1]

	command, exists := allCommands[describeCommand]

	commandPath, isInPath := SearchCommandPath(describeCommand)

	if exists {
		fmt.Println(command.Description)

	} else if isInPath {
		fmt.Printf("%v is %v\n", describeCommand, commandPath)
	} else {
		fmt.Println(describeCommand + ": not found")
	}
}

func GetCommand(command string) (Command, bool) {
	value, isExist := allCommands[command]

	return value, isExist
}

func init() {
	registerCommands()
}

func SearchCommandPath(command string) (string, bool) {
	path := os.Getenv("PATH")
	pathDirs := strings.Split(path, string(os.PathListSeparator))

	for _, dir := range pathDirs {
		cmdPath := dir + string(os.PathSeparator) + command

		_, err := os.Stat(cmdPath)

		if err == nil {
			return cmdPath, true
		}
	}

	return "", false
}
