package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commandUtils "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands/concrete"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
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
		//userInput := strings.Split(TrimNewLine(rawInput), " ")

		params := getParameters(trimmedInput)
		userCommand := params[0]

		if err != nil {
			fmt.Println("Error reading input: " + err.Error())
			os.Exit(1)
		}

		command := getCommand(userCommand, params[1:])
		invoker.ExecuteCommand(command)
	}
}

// func retrieveMetadata(commandName string) commandUtils.CommandMetadata {
// 	switch commandName {
// 	case "echo":
// 		return Echo
// 	}
// }

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

// ' ' should group params together
func getParameters(userInput string) (params []string) {
	var current string
	inSingleQuotes := false
	inDoubleQuotes := false

	for _, char := range userInput {
		switch char {
		case '\'':
			if !inDoubleQuotes {
				inSingleQuotes = !inSingleQuotes
			}
		case '"':
			fmt.Println("Double Quotes")
			if !inSingleQuotes {
				inDoubleQuotes = !inDoubleQuotes
			}

		case ' ':
			if !inSingleQuotes && !inDoubleQuotes {
				if current != "" {
					params = append(params, current)
					current = ""
				}
			} else {
				current += string(char)
			}
		default:
			current += string(char)
		}
	}

	if current != "" {
		params = append(params, current)
	}

	return params
}

func TrimNewLine(prompt string) string {
	return strings.TrimRight(prompt, "\r\n")
}
