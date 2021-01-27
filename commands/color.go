package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/matthewlowe/RoboJoshGo/framework"
)

const (
	baseUrl string = "https://singlecolorimage.com/get/"
	width   string = "500"
	height  string = "500"
)

func ColorCommand(context *framework.Context) error {
	fmt.Println("nice")
	code := context.Args[1]

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

	_, err := context.Session.ChannelMessageSendEmbed(context.Channel.ID, &embed)

	return err
}
