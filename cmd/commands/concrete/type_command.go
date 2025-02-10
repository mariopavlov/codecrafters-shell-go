package commands

import commands "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"

type TypeCommand struct {
	command  string
	receiver *commands.TypeReceiver
}

func (tc TypeCommand) Execute() {
	tc.receiver.DescribeCommandOnPath(tc.command)
}

func NewTypeCommand(commandToFind string, receiver *commands.TypeReceiver) TypeCommand {
	return TypeCommand{
		command:  commandToFind,
		receiver: receiver,
	}
}
