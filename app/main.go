package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	name    string
	details string
}

type RegisteredCommands struct {
	commands map[string]*Command
}

func (rc *RegisteredCommands) registerCommand(command string, description string) {
	// Known limitation I'm not checking if command exists already
	rc.commands[command] = &Command{
		name:    command,
		details: description,
	}
}

func (rc *RegisteredCommands) getCommand(command string) (*Command, error) {
	cmd, ok := rc.commands[command]

	if !ok {
		return nil, fmt.Errorf("%s: not found", command)
	}

	return cmd, nil
}

type Shell struct {
	rc *RegisteredCommands
}

func NewShell() *Shell {
	builtInCommands := &RegisteredCommands{
		commands: make(map[string]*Command),
	}
	builtInCommands.registerCommand("echo", "prints string back")
	builtInCommands.registerCommand("type", "prints whether a command is built-in")
	builtInCommands.registerCommand("exit", "exits eval and closes the shell")

	return &Shell{
		rc: builtInCommands,
	}
}

func (s *Shell) typeCommand(command string) {
	cmd, err := s.rc.getCommand(command)
	if err == nil {
		fmt.Printf("%s is a shell builtin\n", cmd.name)
		return
	}

	fullPath, err := exec.LookPath(command)
	if err == nil {
		fmt.Printf("%s is %s\n", command, fullPath)
		return
	}

	fmt.Printf("%s: not found\n", command)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	shell := NewShell()

	for {
		command, args := readInput(reader)
		executeCommand(shell, command, args)
	}
}

func readInput(reader *bufio.Reader) (string, []string) {
	fmt.Print("$ ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	input = strings.TrimSpace(input)
	commands := strings.Split(input, " ")

	return commands[0], commands[1:]
}

func executeCommand(shell *Shell, command string, args []string) {
	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Println(strings.Join(args, " "))
	case "type":
		shell.typeCommand(args[0])
	default:
		tryExecuteCommand(command, args)
	}
}

func tryExecuteCommand(command string, args []string) {
	_, err := isCommand(command)
	if err == nil {
		result, err := exec.Command(command, args...).Output()
		if err == nil {
			fmt.Print(string(result))
			return
		}

		fmt.Println(err)
	}

	fmt.Println(err)
}

func isCommand(command string) (string, error) {
	fullPath, err := exec.LookPath(command)
	if err == nil {
		return fullPath, nil
	}

	return "", fmt.Errorf("%s: command not found", command)
}
