package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/matthewlowe/Robojosh/framework"
)

func InfoCommand(context *framework.Context) error {
	if len(context.Args) <= 1 {
		_, err := context.Session.ChannelMessageSend(context.TextChannel.ID, "Usage: "+context.Prefix+"info @<user>")
		return err
	}

	if context.Args[1][:2] != "<@" || context.Args[1][len(context.Args[1])-1:] != ">" {
		_, err := context.Session.ChannelMessageSend(context.TextChannel.ID, "Usage: "+context.Prefix+"info @<user>")
		return err
	}

	user, err := context.Session.User(framework.TagToUserId(context.Args[1]))

	if err != nil {
		context.Session.ChannelMessageSend(context.TextChannel.ID, "Invalid user, dufus!")
	}

	fields := []*discordgo.MessageEmbedField{
		&discordgo.MessageEmbedField{
			Name:   "ID",
			Value:  user.ID,
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Nickname",
			Value:  user.Username,
			Inline: true,
		},
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    "brought to you by your local electronic nugget lover",
		IconURL: context.Session.State.User.AvatarURL("256x256"),
	}

	embed := discordgo.MessageEmbed{
		Type:        discordgo.EmbedType("rich"),
		Title:       "User info",
		Description: user.Username + "#" + user.Discriminator,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:    user.AvatarURL("256x256"),
			Width:  256,
			Height: 256,
		},
		Color:  10038562, // DARKER_RED
		Footer: &footer,
		Fields: fields,
	}

	_, err = context.Session.ChannelMessageSendEmbed(context.TextChannel.ID, &embed)

	return err
}
