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

func ExecuteExternalCommand(command string, params []string) {
	if ExternalCommandExists(command) {
		utils.ExecuteExternalCommand(command, params)
	} else {
		fmt.Println(command + ": command not found")
	}
}

func ExternalCommandExists(command string) bool {
	if _, exists := utils.SearchCommandPath(command); exists {
		return true
	}

	return false
}
