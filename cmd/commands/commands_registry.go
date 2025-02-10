package commands

type CommandFactory func(args []string) Command

type CommandsRegistry struct {
	factories map[string]CommandFactory
}

func NewCommandsRegistry() *CommandsRegistry {
	return &CommandsRegistry{
		factories: make(map[string]CommandFactory),
	}
}

func (r *CommandsRegistry) Register(name string, factory CommandFactory) {
	r.factories[name] = factory
}

func (r *CommandsRegistry) CreateCommand(name string, args []string) (Command, bool) {
	factory, exists := r.factories[name]
	if !exists {
		return nil, false
	}

	return factory(args), true
}

// type CommandsRegistry struct {
// 	var allCommands = make(map[string] BuiltinCommand)
// }

// var allCommands = make(map[string]Command)

// func registerCommands() {

// func handleChangeDirectory(args []string) {
// 	if len(args) <= 1 {
// 		return
// 	}

// 	directory := args[1]

// 	command := commands.NewChangeDirectoryCommand(directory)
// 	command.Execute()
// }

// func handlePwd(args []string) {
// 	command := commands.NewPwdCommand()
// 	command.Execute()
// }

// func handleExit(args []string) {
// 	command := commands.NewExitCommand()
// 	command.Execute()
// }

// func handleEcho(args []string) {
// 	var message string = strings.Join(args[1:], " ")

// 	command := commands.NewEchoCommand(message)
// 	command.Execute()

// }

// func handleType(args []string) {
// 	if len(args) <= 1 {
// 		fmt.Println("Error: Type command requires command as a parameter")
// 	}

// 	commandToDescribe := args[1]

// 	command := commands.NewTypeCommand(commandToDescribe)

// 	internalCommand, exists := allCommands[commandToDescribe]

// 	if exists {
// 		fmt.Println(internalCommand.Description)
// 	} else {
// 		command.Execute()
// 	}
// }

// func GetCommand(command string) (Command, bool) {
// 	value, isExist := allCommands[command]

// 	return value, isExist
// }

// func init() {
// 	registerCommands()
// }
