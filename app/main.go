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
	command := readCommand()

	switch command {
	case "cd":
	default:
		fmt.Printf("%s: command not found\n", command)
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
