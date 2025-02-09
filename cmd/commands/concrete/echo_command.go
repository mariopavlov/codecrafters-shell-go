package concrete

import "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"

type EchoCommand struct {
	message  string
	receiver receivers.EchoReceiver
}

func (ec EchoCommand) Execute() {
	ec.receiver.PrintToConsole(ec.message)
}

func NewEchoCommand(message string) EchoCommand {
	return EchoCommand{
		message:  message,
		receiver: receivers.EchoReceiver{},
	}
}
