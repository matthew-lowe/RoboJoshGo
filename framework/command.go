package framework

type (
	Command func(ctx *Context) error

	CommandStruct struct {
		Command Command
		Help    string
		Usage   string
	}

	CommandMap map[string]CommandStruct

	CommandRegistry struct {
		commands CommandMap
	}
)

func (handler *CommandRegistry) Register(name, help, usage string, command Command) {
	commandStruct := CommandStruct{
		Command: command,
		Help:    help,
		Usage:   usage,
	}

	handler.commands[name] = commandStruct
}

func (handler *CommandRegistry) Get(name string) (*Command, bool) {
	command, ok := handler.commands[name]

	if ok {
		return &command.Command, true
	} else {
		return nil, false
	}
}

func (handler *CommandRegistry) GetCommandMap() *CommandMap {
	return &handler.commands
}

func NewCommandHandler() *CommandRegistry {
	return &CommandRegistry{commands: make(CommandMap)}
}
