package commands

import "fmt"

type EchoReceiver struct {
}

func (er EchoReceiver) PrintToConsole(message string) {
	fmt.Println(message)
}
