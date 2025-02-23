package commands

import (
	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
)

type TypeCommand struct {
	describeCommand *commands.CommandMetadata
	receiver        *receivers.TypeReceiver
	metadata        commands.CommandMetadata
}

func (tc *TypeCommand) Execute() {
	// TODO Fix, it should get metadata from the command
	tc.receiver.DescribeCommand(*tc.describeCommand)
}

func GetTypeMetadata() commands.CommandMetadata {
	return commands.NewCommandMetadata(
		"type",
		"type is a shell builtin",
		"type [command]",
	)
}

func (tc *TypeCommand) Metadata() commands.CommandMetadata {
	return tc.metadata
}

func NewTypeCommand(command *commands.CommandMetadata, receiver *receivers.TypeReceiver) *TypeCommand {
	return &TypeCommand{
		describeCommand: command,
		receiver:        receiver,
		metadata: commands.NewCommandMetadata(
			"type",
			"type is a shell builtin",
			"type [command]",
		),
	}
}
