package commands

import "fmt"

type Invoker struct {
	Command string
}

func NewInvoker() *Invoker {
	return &Invoker{}
}

func (i *Invoker) ExecuteCommand(command Command) {
	fmt.Printf("Execute Command: %v\n", command)
	command.Execute()
}
