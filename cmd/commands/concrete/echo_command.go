package commands

import commands "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"

type EchoCommand struct {
	message  string
	receiver commands.EchoReceiver
}

func (ec EchoCommand) Execute() {
	ec.receiver.PrintToConsole(ec.message)
}

func NewEchoCommand(echoMessage string) EchoCommand {
	return EchoCommand{
		message:  echoMessage,
		receiver: commands.EchoReceiver{},
	}
}
