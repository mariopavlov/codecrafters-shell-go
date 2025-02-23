package commands

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/cmd/utils"
)

type ExternalReceiver struct {
}

func NewExternalReceiver() *ExternalReceiver {
	return &ExternalReceiver{}
}

func (er ExternalReceiver) ExternalCommandExists(command string) (string, bool) {
	if path, exists := utils.SearchCommandPath(command); exists {
		return path, true
	}

	return "", false
}

func (er ExternalReceiver) ExecuteExternalCommand(externalCommand string, args []string) {
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
