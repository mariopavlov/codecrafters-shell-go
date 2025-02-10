package commands

import (
	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
)

type TypeCommand struct {
	command  string
	receiver *receivers.TypeReceiver
	metadata commands.CommandMetadata
}

func (tc *TypeCommand) Execute() {
	tc.receiver.DescribeCommandOnPath(tc.command)
}

func (tc *TypeCommand) Metadata() commands.CommandMetadata {
	return tc.metadata
}

func NewTypeCommand(commandToFind string, receiver *receivers.TypeReceiver) *TypeCommand {
	return &TypeCommand{
		command:  commandToFind,
		receiver: receiver,
		metadata: commands.NewCommandMetadata(
			"type",
			"type is a shell builtin",
			"type [command]",
		),
	}
}
