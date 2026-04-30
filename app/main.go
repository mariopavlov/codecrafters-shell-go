package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		command, args := readCommand(reader)
		executeCommand(command, args)
	}
}

func readCommand(reader *bufio.Reader) (string, []string) {
	fmt.Print("$ ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	input = strings.TrimSpace(input)
	commands := strings.Split(input, " ")

	return commands[0], commands[1:]
}

func executeCommand(command string, args []string) {
	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Println(strings.Join(args, " "))
	default:
		fmt.Printf("%s: command not found\n", command)
	}
}
