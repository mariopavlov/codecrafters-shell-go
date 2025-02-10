package commands

import (
	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
)

type EchoCommand struct {
	message  string
	receiver *receivers.EchoReceiver
	metadata commands.CommandMetadata
}

func (ec *EchoCommand) Execute() {
	ec.receiver.PrintToConsole(ec.message)
}

func (ec *EchoCommand) Metadata() commands.CommandMetadata {
	return ec.metadata
}

func NewEchoCommand(echoMessage string, receiver *receivers.EchoReceiver) *EchoCommand {
	return &EchoCommand{
		message:  echoMessage,
		receiver: receiver,
		metadata: commands.NewCommandMetadata(
			"echo",
			"echo is a shell builtin",
			"echo [text to print out]",
		),
	}
}
