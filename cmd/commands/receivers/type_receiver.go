package commands

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/cmd/utils"
)

type TypeReceiver struct {
}

func (tr TypeReceiver) DescribeCommandOnPath(command string) {
	commandPath, exists := utils.SearchCommandPath(command)

	if exists {
		fmt.Printf("%v is %v\n", command, commandPath)
	} else {
		fmt.Println(command + ": not found")
	}
}
