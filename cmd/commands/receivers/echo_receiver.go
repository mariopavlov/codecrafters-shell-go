package commands

import "fmt"

type EchoReceiver struct {
}

func NewEchoReceiver() *EchoReceiver {
	return &EchoReceiver{}
}

func (er EchoReceiver) PrintToConsole(message string) {
	fmt.Println(message)
}
