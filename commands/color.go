package commands

import (
	"github.com/matthew-lowe/RoboJoshGo/framework"
)

const (
	baseUrl string = "https://singlecolorimage.com/get/"
	width   string = "500"
	height  string = "500"
)

func ColorCommand(context *framework.Context) error {
	/*
		code := "#FFFFFF" //context.Args[1]

		valid, err := framework.VerifyHexColor(code)

		if err != nil {
			return err
		}

		if !valid {
			field := discordgo.MessageEmbedField{
				Name:  "You fucking idiot",
				Value: "Invalid color code provided! Must be in form #<code> where <code> is 6 hexadecimal digits",
			}

			return context.ReplyRichEmbed("Invalid hex code", "", []*discordgo.MessageEmbedField{&field})
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

		return context.Reply("sadge") //ReplyImageEmbed("#"+code, "Requested by "+context.User.Username+"#"+context.User.Discriminator, &image)
	*/
	return context.Reply("sadge")
}
