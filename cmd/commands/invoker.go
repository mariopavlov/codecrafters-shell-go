package commands

type Invoker struct {
	Command string
}

func NewInvoker() Invoker {
	return Invoker{}
}

func (i Invoker) ExecuteCommand(command Command) {
	command.Execute()
}
