package framework

import (
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Discord     *discordgo.Session
	User        *discordgo.User
	TextChannel *discordgo.Channel
	Guild       *discordgo.Guild
	Message     *discordgo.Message
	Args        []string
	Prefix      string

	CmdRegistry *CommandRegistry
}
