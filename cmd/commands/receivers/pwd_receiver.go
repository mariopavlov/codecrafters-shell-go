package commands

import (
	"fmt"
	"os"
)

type PwdReceiver struct {
}

func (pc PwdReceiver) PrintWorkingDirectory() {
	currentDirectory, error := os.Getwd()

	if error == nil {
		fmt.Println(currentDirectory)
	}
}
