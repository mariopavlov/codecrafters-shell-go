package commands

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/cmd/commands"
	"github.com/codecrafters-io/shell-starter-go/cmd/utils"
)

type TypeReceiver struct {
	registry *commands.CommandsRegistry
}

func NewTypeReceiver(registry *commands.CommandsRegistry) *TypeReceiver {
	return &TypeReceiver{
		registry: registry,
	}
}

func (tr TypeReceiver) DescribeCommandOnPath(command string) {
	if builtin, exists := tr.registry.CreateCommand(command, []string{}); exists {
		fmt.Println(builtin.Metadata().Description)
		return
	}

	commandPath, exists := utils.SearchCommandPath(command)

	if exists {
		fmt.Printf("%v is %v\n", command, commandPath)
	} else {
		fmt.Println(command + ": not found")
	}
}
