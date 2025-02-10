package commands

import (
	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
)

type PwdCommand struct {
	receiver *receivers.DirectoryReceiver
	metadata commands.CommandMetadata
}

func (pc *PwdCommand) Execute() {
	pc.receiver.PrintWorkingDirectory()
}

func (pc *PwdCommand) Metadata() commands.CommandMetadata {
	return pc.metadata
}

func NewPwdCommand(receiver *receivers.DirectoryReceiver) PwdCommand {
	return PwdCommand{
		receiver: receiver,
		metadata: commands.NewCommandMetadata(
			"pwd",
			"pwd is a shell builtin",
			"pwd [command]",
		),
	}
}
