package commands

import commands "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"

type ChangeDirectoryCommand struct {
	directory string
	receiver  *commands.DirectoryReceiver
}

func (cd ChangeDirectoryCommand) Execute() {
	cd.receiver.ChangeDirectory(cd.directory)
}

func NewChangeDirectoryCommand(newDirectory string, receiver *commands.DirectoryReceiver) ChangeDirectoryCommand {
	return ChangeDirectoryCommand{
		directory: newDirectory,
		receiver:  receiver,
	}
}
