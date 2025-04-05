package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commandUtils "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands/concrete"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
	"github.com/codecrafters-io/shell-starter-go/cmd/utils"
)

func main() {
	// Get to Basics
	// Transfer here all commands
	// Invoker
	invoker := commandUtils.NewInvoker()

	// All Receivers
	// directoryReceiver := receivers.NewDirectoryReceiver()

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		trimmedInput := TrimNewLine((rawInput))
		// userInput := strings.Split(TrimNewLine(rawInput), " ")

		params := utils.ParseArguments(trimmedInput)
		userCommand := params[0]

		if err != nil {
			fmt.Println("Error reading input: " + err.Error())
			os.Exit(1)
		}

		command := getCommand(userCommand, params[1:])
		invoker.ExecuteCommand(command)
	}
}

func getCommand(command string, params []string) commandUtils.Command {
	echoReceiver := receivers.NewEchoReceiver()
	exitReceiver := receivers.NewExitReceiver()
	typeReceiver := receivers.NewTypeReceiver()
	directoryReceiver := receivers.NewDirectoryReceiver()
	externalReceiver := receivers.NewExternalReceiver()

	switch command {
	case "echo":
		// Build Echo
		return commands.NewEchoCommand(strings.Join(params, " "), echoReceiver)
	case "exit":
		// Build Exit
		return commands.NewExitCommand(exitReceiver)
	case "type":
		// Build Type
		var commandMetadata commandUtils.CommandMetadata
		if params[0] == "type" {
			commandMetadata = commands.GetTypeMetadata()
		} else {
			commandMetadata = getCommand(params[0], params).Metadata()
		}

		return commands.NewTypeCommand(&commandMetadata, typeReceiver)
	case "pwd":
		// Build PWD
		return commands.NewPwdCommand(directoryReceiver)
	case "cd":
		return commands.NewChangeDirectoryCommand(params[0], directoryReceiver)
	default:
		return commands.NewExternalCommand(command, params, externalReceiver)
	}
}

func TrimNewLine(prompt string) string {
	return strings.TrimRight(prompt, "\r\n")
}
