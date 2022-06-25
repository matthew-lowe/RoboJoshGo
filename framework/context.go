package framework

import (
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Session     *discordgo.Session
	User        *discordgo.User
	Channel     *discordgo.Channel
	Guild       *discordgo.Guild
	Interaction *discordgo.InteractionCreate
}

func (context *Context) Reply(message string) error {
	err := context.Session.InteractionRespond(context.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "slash commands poggers verified",
		},
	})

	return err
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

	err := context.Session.InteractionRespond(context.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&embed},
		},
	})

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

	err := context.Session.InteractionRespond(context.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&embed},
		},
	})

	return err
}

func (context *Context) ReplyFromEmbed(embed *discordgo.MessageEmbed) error {
	err := context.Session.InteractionRespond(context.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	return err
}
