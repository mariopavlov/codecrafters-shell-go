package commands

import "fmt"

type EchoReceiver struct {
}

func (p EchoReceiver) PrintToConsole(message string) {
	fmt.Println(message)
}
