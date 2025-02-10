package commands

import (
	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
)

type ExitCommand struct {
	receiver *receivers.ExitReceiver
	metadata commands.CommandMetadata
}

func (ex *ExitCommand) Execute() {
	ex.receiver.DefaultExit()
}

func (ex *ExitCommand) Metadata() commands.CommandMetadata {
	return ex.metadata
}

func NewExitCommand(receiver *receivers.ExitReceiver) ExitCommand {
	return ExitCommand{
		receiver: receiver,
		metadata: commands.NewCommandMetadata(
			"exit",
			"exit is a shell builtin",
			"exit [code]",
		),
	}
}
