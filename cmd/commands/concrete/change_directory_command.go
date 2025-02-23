package commands

import (
	commands "github.com/codecrafters-io/shell-starter-go/cmd/commands"
	receivers "github.com/codecrafters-io/shell-starter-go/cmd/commands/receivers"
)

type ChangeDirectoryCommand struct {
	directory string
	receiver  *receivers.DirectoryReceiver
	metadata  commands.CommandMetadata
}

func (cd ChangeDirectoryCommand) Execute() {
	cd.receiver.ChangeDirectory(cd.directory)
}

func (cd ChangeDirectoryCommand) Metadata() commands.CommandMetadata {
	return cd.metadata
}

func NewChangeDirectoryCommand(newDirectory string, receiver *receivers.DirectoryReceiver) *ChangeDirectoryCommand {
	return &ChangeDirectoryCommand{
		directory: newDirectory,
		receiver:  receiver,
		metadata: commands.NewCommandMetadata(
			"cd",
			"Change current working directory",
			"cd [directory]",
		),
	}
}
