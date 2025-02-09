package commands

import (
	"fmt"
	"os"
	"strings"
)

type TypeReceiver struct {
}

func (tr TypeReceiver) DescribeCommandOnPath(command string) {
	commandPath, exists := searchCommandPath(command)

	if exists {
		fmt.Printf("%v is %v\n", command, commandPath)
	} else {
		fmt.Println(command + ": not found")
	}
}

func searchCommandPath(command string) (string, bool) {
	path := os.Getenv("PATH")
	pathDirs := strings.Split(path, string(os.PathListSeparator))

	for _, dir := range pathDirs {
		cmdPath := dir + string(os.PathSeparator) + command

		_, err := os.Stat(cmdPath)

		if err == nil {
			return cmdPath, true
		}
	}

	return "", false
}
