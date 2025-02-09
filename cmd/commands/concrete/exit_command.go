package commands

import (
	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
)

type ExitCommand struct {
	receiver commands.ExitReceiver
}

func (ex ExitCommand) Execute() {
	ex.receiver.DefaultExit()
}

func NewExitCommand() ExitCommand {
	return ExitCommand{
		receiver: commands.ExitReceiver{},
	}
}
