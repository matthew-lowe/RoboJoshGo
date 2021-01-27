package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/matthewlowe/RoboJoshGo/framework"
	"strconv"
)

func SnipeCommand(context *framework.Context) error {
	var msg *discordgo.Message

	if len(context.Args) <= 1 {
		msg = context.DeletionHistory.GetLastMessage(0)

	} else {
		amt, err := strconv.Atoi(context.Args[1])

		if err != nil {
			context.Reply("Usage: snipe [index]")
		}

		msg = context.DeletionHistory.GetLastMessage(amt)
	}

	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Username",
			Value:  msg.Author.Username,
			Inline: false,
		},
		{
			Name:   "Message",
			Value:  msg.Content,
			Inline: false,
		},
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    "brought to you by your local electronic nugget lover",
		IconURL: context.Session.State.User.AvatarURL("256x256"),
	}

	embed := discordgo.MessageEmbed{
		Type:   discordgo.EmbedType("rich"),
		Title:  "Deletion Snipe",
		Color:  10038562, // DARKER_RED
		Footer: &footer,
		Fields: fields,
	}

	context.Session.ChannelMessageSendEmbed(context.Channel.ID, &embed)

	return nil
}
