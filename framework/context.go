package framework

import (
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Session     *discordgo.Session
	User        *discordgo.User
	TextChannel *discordgo.Channel
	Guild       *discordgo.Guild
	Message     *discordgo.Message
	Args        []string
	Prefix      string

	CmdRegistry *CommandRegistry
}
