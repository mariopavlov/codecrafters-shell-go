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
	command := exec.Command(externalCommand, args...)

	var outBuffer bytes.Buffer
	command.Stdout = &outBuffer
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		fmt.Printf("%v: command not found\n", externalCommand)
	}

	fmt.Print(outBuffer.String())
}
