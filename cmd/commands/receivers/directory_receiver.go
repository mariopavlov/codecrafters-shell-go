package commands

import (
	"fmt"
	"os"
)

type DirectoryReceiver struct {
}

func NewDirectoryReceiver() *DirectoryReceiver {
	return &DirectoryReceiver{}
}

func (dr DirectoryReceiver) PrintWorkingDirectory() {
	currentDirectory, error := os.Getwd()

	if error == nil {
		fmt.Println(currentDirectory)
	}
}

func (dr DirectoryReceiver) ChangeDirectory(directory string) {
	if len(directory) == 0 {
		return
	}

	if directory == "~" {
		homeDirectory, _ := os.UserHomeDir()
		os.Chdir(homeDirectory)
		return
	}

	error := os.Chdir(directory)

	if error != nil {
		fmt.Printf("cd: %v: No such file or directory\n", directory)
		return
	}
}
