package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/matthewlowe/RoboJoshGo/framework"
)

const (
	baseUrl string = "https://singlecolorimage.com/get/"
	width   string = "500"
	height  string = "500"
)

func ColorCommand(context *framework.Context) error {
	code := context.Args[1]

	valid, err := framework.VerifyHexColor(code)

	if err != nil {
		return err
	}

	if !valid {
		footer := discordgo.MessageEmbedFooter{
			Text:    "brought to you by your local electronic nugger lover",
			IconURL: context.Session.State.User.AvatarURL("256x256"),
		}

		field := discordgo.MessageEmbedField{
			Name:  "You fucking idiot",
			Value: "Invalid color code provided! Must be in form #<code> where <code> is 6 hexadecimal digits",
		}

		embed := discordgo.MessageEmbed{
			Type:   discordgo.EmbedType("rich"),
			Title:  "Invalid hex code!",
			Color:  10038562, // DARKER_RED
			Fields: []*discordgo.MessageEmbedField{&field},
			Footer: &footer,
		}

		_, err := context.Session.ChannelMessageSendEmbed(context.Channel.ID, &embed)

		return err
	}

	if code[0] == '#' {
		code = code[1:]
	}

	url := baseUrl + code + "/" + width + "x" + height

	image := discordgo.MessageEmbedImage{
		URL:    url,
		Width:  500,
		Height: 500,
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    "brought to you by your local electronic nugget lover",
		IconURL: context.Session.State.User.AvatarURL("256x256"),
	}

	embed := discordgo.MessageEmbed{
		URL:         url,
		Type:        "image",
		Title:       "#" + code,
		Description: "Requested by " + context.User.Username + "#" + context.User.Discriminator,
		Image:       &image,
		Color:       10038562, // DARKER_RED
		Footer:      &footer,
	}

	_, err = context.Session.ChannelMessageSendEmbed(context.Channel.ID, &embed)

	return err
}
