package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		userInput := strings.Split(TrimNewLine(rawInput), " ")

		if err != nil {
			fmt.Println("Error reading input: " + err.Error())
			os.Exit(1)
		}

		// Get Command
		command, exists := GetCommand(userInput[0])
		if exists {
			command.Handler(userInput)
		} else {
			fmt.Println(command.Name + ": command not found")
		}
	}

}

func TrimNewLine(prompt string) string {
	return strings.TrimRight(prompt, "\r\n")
}