package commands

import commands "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"

type PwdCommand struct {
	receiver commands.DirectoryReceiver
}

func (pc PwdCommand) Execute() {
	pc.receiver.PrintWorkingDirectory()
}

func NewPwdCommand() PwdCommand {
	return PwdCommand{
		receiver: commands.DirectoryReceiver{},
	}
}
