package main

import (
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Handler     func(args []string) error
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
	commands["pwd"] = Command{
		Name:        "pwd",
		Description: "pwd is a shell builtin",
		Usage:       "pwd [command]",
		Handler:     commandPwd,
	}
	commands["cd"] = Command{
		Name:        "cd",
		Description: "Change current working directory",
		Usage:       "cd [directory]",
		Handler:     commandDirectoryChange,
	}
}

func commandDirectoryChange(args []string) error {
	if len(args) <= 1 {
		return nil
	}

	if args[1] == "~" {
		// TODO Do we have case to fail? What is the case that can fail here?
		homeDirectory, _ := os.UserHomeDir()

		return os.Chdir(homeDirectory)
	}

	error := os.Chdir(args[1])

	if error != nil {
		fmt.Printf("cd: %v: No such file or directory\n", args[1])
	}

	return nil
}

// TODO This method does not need args, but need to satisfy the command handler interface
// improve this code
func commandPwd(args []string) error {
	currentDir, error := os.Getwd()

	if error == nil {
		fmt.Println(currentDir)
	}

	return nil
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

	commandPath, isInPath := SearchCommandPath(describeCommand)

	if exists {
		fmt.Println(command.Description)

	} else if isInPath {
		fmt.Printf("%v is %v\n", describeCommand, commandPath)
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
