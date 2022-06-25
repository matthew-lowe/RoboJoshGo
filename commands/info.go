package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/matthew-lowe/RoboJoshGo/framework"
)

func InfoCommand(context *framework.Context) error {
	userId := fmt.Sprintf("%s", context.Interaction.ApplicationCommandData().Options[0].Value)

	user, err := context.Session.User(userId)

	if err != nil {
		_, err = context.Session.ChannelMessageSend(context.Channel.ID, "Invalid user, dufus!")

		return err
	}

	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "ID",
			Value:  user.ID,
			Inline: true,
		},
		{
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

	return context.ReplyFromEmbed(&embed)
}
