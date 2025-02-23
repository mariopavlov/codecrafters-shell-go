package commands

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/cmd/utils"
)

type ExternalReceiver struct {
}

func NewExternalReceiver() *ExternalReceiver {
	return &ExternalReceiver{}
}

func (er ExternalReceiver) ExecuteExternalCommand(command string, params []string) {
	if _, exists := utils.SearchCommandPath(command); exists {
		utils.ExecuteExternalCommand(command, params)
	} else {
		fmt.Println(command + ": command not found")
	}
}

func (er ExternalReceiver) ExternalCommandExists(command string) (string, bool) {
	if path, exists := utils.SearchCommandPath(command); exists {
		return path, true
	}

	return "", false
}
