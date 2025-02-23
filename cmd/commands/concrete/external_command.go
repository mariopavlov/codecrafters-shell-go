package commands

import (
	"github.com/codecrafters-io/shell-starter-go/cmd/commands"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
)

type ExternalCommand struct {
	command  string
	params   []string
	receiver *receivers.ExternalReceiver
}

func (ec ExternalCommand) Metadata() commands.CommandMetadata {
	// External Command should be in the form %COMMAND is %PATH
	return commands.NewCommandMetadata(
		ec.command,
		"External Command",
		"External Command Help Placeholder",
	)
}

func NewExternalCommand(command string,
	params []string,
	receiver *receivers.ExternalReceiver) *ExternalCommand {
	return &ExternalCommand{
		command:  command,
		params:   params,
		receiver: receiver,
	}
}
