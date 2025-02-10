package commands

type Command interface {
	Execute()
	Metadata() CommandMetadata
}
