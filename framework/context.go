package framework

import (
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Session *discordgo.Session
	User    *discordgo.User
	Channel *discordgo.Channel
	Guild   *discordgo.Guild
	Message *discordgo.Message
	Args    []string
	Prefix  string

	CmdRegistry *CommandRegistry
}

func (context *Context) Reply(message string) (*discordgo.Message, error) {
	msg, err := context.Session.ChannelMessageSend(context.Channel.ID, context.User.Mention()+" "+message)

	return msg, err
}
