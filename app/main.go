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
	for {
		command := readCommand()
		executeCommand(command)
	}
}

func readCommand() string {
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	return strings.TrimSpace(command)
}

func executeCommand(command string) {
	switch command {
	case "cd":
	default:
		fmt.Printf("%s: command not found\n", command)
	}
}
