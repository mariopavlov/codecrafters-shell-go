package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	concreteCommands "github.com/codecrafters-io/shell-starter-go/cmd/commands/concrete"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
	"github.com/codecrafters-io/shell-starter-go/cmd/utils"
)

func main() {
	registry := commands.NewCommandsRegistry()
	registerCommands(registry)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		userInput := strings.Split(TrimNewLine(rawInput), " ")
		userCommand := userInput[0]

		if err != nil {
			fmt.Println("Error reading input: " + err.Error())
			os.Exit(1)
		}

		command, exists := registry.CreateCommand(userCommand, userInput[1:])

		if exists {
			command.Execute()
		} else if _, exists := utils.SearchCommandPath(userCommand); exists {
			ExecuteExternalCommand(userCommand, userInput[1:])
		} else {
			fmt.Println(userInput[0] + ": command not found")
		}
	}
}

func registerCommands(registry *commands.CommandsRegistry) {
	registry.Register("type", func(args []string) commands.Command {
		if len(args) < 1 {
			return nil
		}

		return concreteCommands.NewTypeCommand(args[0], receivers.NewTypeReceiver(registry))
	})

	registry.Register("echo", func(args []string) commands.Command {
		return concreteCommands.NewEchoCommand(args[0], receivers.NewEchoReceiver())
	})
}

func ExecuteExternalCommand(externalCommand string, args []string) {
	var command *exec.Cmd

	if len(args) >= 1 {
		command = exec.Command(externalCommand, args...)
	} else {
		command = exec.Command(externalCommand, args...)

	}

	var outBuffer bytes.Buffer
	command.Stdout = &outBuffer
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		fmt.Println("Error executing command:", err)
	}

	fmt.Print(outBuffer.String())
}

func TrimNewLine(prompt string) string {
	return strings.TrimRight(prompt, "\r\n")
}
