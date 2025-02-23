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
	echoReceiver := receivers.NewEchoReceiver()
	exitReceiver := receivers.NewExitReceiver()
	typeReceiver := receivers.NewTypeReceiver()
	directoryReceiver := receivers.NewDirectoryReceiver()

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
			echoCommand := commands.NewEchoCommand(strings.Join(userInput[1:], " "), echoReceiver)
			invoker.ExecuteCommand(echoCommand)
		case "exit":
			// Build Exit
			exitCommand := commands.NewExitCommand(exitReceiver)
			invoker.ExecuteCommand(exitCommand)
		case "type":
			// Build Type
			// describe := retrieveMetadata(userInput[1])
			typeCommand := commands.NewTypeCommand(userInput[1], typeReceiver)
			invoker.ExecuteCommand(typeCommand)
		case "pwd":
			// Build PWD
			pwdCommand := commands.NewPwdCommand(directoryReceiver)
			invoker.ExecuteCommand(pwdCommand)
		case "cd":
			changeDirectoryCommand := commands.NewChangeDirectoryCommand(userInput[1], directoryReceiver)
			invoker.ExecuteCommand(changeDirectoryCommand)

		default:
			if _, exists := utils.SearchCommandPath(userCommand); exists {
				utils.ExecuteExternalCommand(userCommand, userInput[1:])
			} else {
				fmt.Println(userInput[0] + ": command not found")
			}

		}
	}
}

// func retrieveMetadata(commandName string) commandUtils.CommandMetadata {
// 	switch commandName {
// 	case "echo":
// 		return Echo
// 	}
// }

func TrimNewLine(prompt string) string {
	return strings.TrimRight(prompt, "\r\n")
}
