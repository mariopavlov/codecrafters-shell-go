package commands

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/cmd/commands"
	"github.com/codecrafters-io/shell-starter-go/cmd/utils"
)

type TypeReceiver struct {
}

func NewTypeReceiver() *TypeReceiver {
	return &TypeReceiver{}
}

func (tr TypeReceiver) DescribeCommand(metadata commands.CommandMetadata) {

	commandPath, exists := utils.SearchCommandPath(metadata.Name)

	if exists {
		fmt.Printf("%v is %v\n", metadata.Name, commandPath)
	} else {
		fmt.Println(metadata.Name + ": not found")
	}
}
