package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

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

		command, exists := GetCommand(userInput[0])
		if exists {
			command.Handler(userInput)
		} else if _, exists := SearchCommandPath(userCommand); exists {
			ExecuteExternalCommand(userCommand, userInput[1:])
		} else {
			fmt.Println(userInput[0] + ": command not found")
		}
	}
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
