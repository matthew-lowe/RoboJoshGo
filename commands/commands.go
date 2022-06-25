package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/matthew-lowe/RoboJoshGo/framework"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping command to see if the bot is working",
		},
		{
			Name:        "color",
			Description: "Generate a color from a hex code",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "hex-code",
					Description: "The hex color-code to generate",
					Required:    true,
				},
			},
		},
		{
			Name:        "eightball",
			Description: "Magic 8 ball!",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "question",
					Description: "The question you want to ask",
					Required:    true,
				},
			},
		},
		{
			Name:        "help",
			Description: "Help command",
		},
		{
			Name:        "info",
			Description: "Get info about a particular user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user you want to dox",
					Required:    true,
				},
			},
		},
	}

	Handlers = map[string]func(*framework.Context) error{
		"ping":      PingCommand,
		"color":     ColorCommand,
		"eightball": EightBallCommand,
		"help":      HelpCommand,
		"info":      InfoCommand,
	}
)
