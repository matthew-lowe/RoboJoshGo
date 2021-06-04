package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/matthewlowe/RoboJoshGo/framework"
)

func HelpCommand(context *framework.Context) error {
	commandMap := *context.CmdRegistry.GetCommandMap()

	names := make([]string, len(commandMap))
	helps := make([]string, len(commandMap))
	usages := make([]string, len(commandMap))
	c := 0

	for name, commandStruct := range commandMap {
		names[c] = name
		helps[c] = commandStruct.Help
		usages[c] = commandStruct.Usage
		c++
	}

	fields := make([]*discordgo.MessageEmbedField, c)

	for i := 0; i < c; i++ {
		fields[i] = &discordgo.MessageEmbedField{
			Name:   names[i],
			Value:  helps[i] + "\nUsage: `" + usages[i] + "`",
			Inline: false,
		}
	}

	return context.ReplyRichEmbed("List of commands: ", "Current prefix: "+context.Prefix, fields)
}
