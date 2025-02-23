package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	concreteCommands "github.com/codecrafters-io/shell-starter-go/cmd/commands/concrete"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
)

func main() {

	// Get to Basics
	// Transfer here all commands

	// All Receivers
	// directoryReceiver := receivers.NewDirectoryReceiver()
	// echoReceiver := receivers.NewEchoReceiver()
	// exitReceiver := receivers.NewExitReceiver()
	// typeReceiver := receivers.NewTypeReceiver()

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

		switch userCommand {
		case "echo":
			// Build Echo
			fmt.Println("Builtin Echo")
		case "exit":
			// Build Exit
			fmt.Println("Builtin Echo")
		case "type":
			// Build Type
			fmt.Println("Builtin Echo")
		case "pwd":
			// Build PWD
			fmt.Println("Builtin Echo")

		}

		// command, exists := registry.CreateCommand(userCommand, userInput[1:])

		// if exists {
		// 	command.Execute()
		// } else if _, exists := utils.SearchCommandPath(userCommand); exists {
		// 	ExecuteExternalCommand(userCommand, userInput[1:])
		// } else {
		// 	fmt.Println(userInput[0] + ": command not found")
		// }
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

func TrimNewLine(prompt string) string {
	return strings.TrimRight(prompt, "\r\n")
}
