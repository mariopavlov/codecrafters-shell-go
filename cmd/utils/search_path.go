package utils

import (
	"os"
	"strings"
)

func SearchCommandPath(command string) (cmdPath string, exists bool) {
	path := os.Getenv("PATH")
	pathDirs := strings.Split(path, string(os.PathListSeparator))

	exists = false

	for _, dir := range pathDirs {
		cmdPath = dir + string(os.PathSeparator) + command

		_, err := os.Stat(cmdPath)

		if err == nil {
			return cmdPath, true
		}
	}

	return "", exists
}
