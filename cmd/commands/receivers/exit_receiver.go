package commands

import "os"

type ExitReceiver struct {
}

func NewExitReceiver(exitCode int) *ExitReceiver {
	return &ExitReceiver{}
}

func (er ExitReceiver) Exit(exitCode int) {
	os.Exit(exitCode)
}
