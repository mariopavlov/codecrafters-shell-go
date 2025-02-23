package commands

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/cmd/commands"
)

type TypeReceiver struct {
}

func NewTypeReceiver() *TypeReceiver {
	return &TypeReceiver{}
}

func (tr TypeReceiver) DescribeCommand(metadata commands.CommandMetadata) {
	fmt.Println(metadata.Description)
}
