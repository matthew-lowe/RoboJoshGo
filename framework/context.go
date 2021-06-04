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

func (context *Context) ReplyRichEmbed(title, description string, fields []*discordgo.MessageEmbedField) error {
	embed := discordgo.MessageEmbed{
		Type:        discordgo.EmbedType("rich"),
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       10038562,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Brought to you by your local electronic nugget lover",
			IconURL: context.Session.State.User.AvatarURL("256x256"),
		},
	}

	_, err := context.Session.ChannelMessageSendEmbed(context.Channel.ID, &embed)

	return err
}

func (context *Context) ReplyImageEmbed(title, description string, image *discordgo.MessageEmbedImage) error {
	embed := discordgo.MessageEmbed{
		Type:        discordgo.EmbedType("image"),
		Title:       title,
		Description: description,
		Color:       10038562,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Brought to you by your local electronic nugget lover",
			IconURL: context.Session.State.User.AvatarURL("256x256"),
		},
		Image: image,
	}

	_, err := context.Session.ChannelMessageSendEmbed(context.Channel.ID, &embed)

	return err
}
