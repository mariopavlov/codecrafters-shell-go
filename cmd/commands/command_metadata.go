package commands

type CommandMetadata struct {
	Name string
	Description string
	Usage string
}

func NewCommandMetadata(name string, description string, usage string) CommandMetadata {
	return CommandMetadata{
		Name: name,
		Description: description,
		Usage: usage,
	}
}