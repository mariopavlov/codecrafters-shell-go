package commands

import "os"

type ExitReceiver struct {
}

func (er ExitReceiver) DefaultExit() {
	os.Exit(0)

}

func (er ExitReceiver) Exit(exitCode int) {
	os.Exit(exitCode)
}
