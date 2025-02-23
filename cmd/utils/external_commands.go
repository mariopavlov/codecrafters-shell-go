package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

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
