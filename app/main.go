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
		command := readCommand(reader)
		executeCommand(command)
	}
}

func readCommand(reader *bufio.Reader) string {
	fmt.Print("$ ")
	command, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	return strings.TrimSpace(command)
}

func executeCommand(command string) {
	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Println("Echo command")
	default:
		fmt.Printf("%s: command not found\n", command)
	}
}
