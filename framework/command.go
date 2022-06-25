package framework

import "github.com/bwmarrin/discordgo"

type (
	CommandHandler func(ctx *Context) error

	CommandMap map[string]*discordgo.ApplicationCommand

	CommandRegistry struct {
		commands CommandMap
	}
)

func (handler *CommandRegistry) Register(command *discordgo.ApplicationCommand, session *discordgo.Session) error {
	newCommand, err := session.ApplicationCommandCreate(session.State.User.ID, "", command)

	if err != nil {
		return err
	}

	handler.commands[command.Name] = newCommand

	return nil
}

func (handler *CommandRegistry) Get(name string) (*discordgo.ApplicationCommand, bool) {
	command, ok := handler.commands[name]

	if ok {
		return command, true
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
