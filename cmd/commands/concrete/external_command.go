package commands

import (
	"fmt"

	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
)

type ExternalCommand struct {
	command  string
	params   []string
	receiver *receivers.ExternalReceiver
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

func (ec *ExternalCommand) Metadata() commands.CommandMetadata {
	// External Command should be in the form %COMMAND is %PATH
	var commandIs string
	if path, exists := ec.receiver.ExternalCommandExists(ec.command); exists {
		commandIs = fmt.Sprintf("%v is %v", ec.command, path)
	} else {
		commandIs = fmt.Sprintf("%v: not found", ec.command)
	}

	return commands.NewCommandMetadata(
		ec.command,
		commandIs,
		"External Command Help Placeholder",
	)
}

func (ec *ExternalCommand) Execute() {
	ec.receiver.ExecuteExternalCommand(ec.command, ec.params)
}
