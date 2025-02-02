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
		} else if commandPath, exists := SearchCommandPath(userCommand); exists {
			ExecuteExternalCommand(commandPath, userInput[1:])
		} else {
			fmt.Println(userInput[0] + ": command not found")
		}
	}

}

func ExecuteExternalCommand(path string, args []string) {
	command := exec.Command(path, args...)

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
