package commands

import (
	"github.com/matthewlowe/Robojosh/framework"
	"math/rand"
)

var (
	options = []string{
		"Yes",
		"No",
		"Perhaps",
		"Likely",
		"Unlikely",
		"Focus harder and ask again",
	}
)

func getRandomResponse() string {
	randIndex := rand.Intn(len(options))

	return options[randIndex]
}

func EightBallCommand(context *framework.Context) error {
	_, err := context.Session.ChannelMessageSend(
		context.TextChannel.ID,
		"<@"+context.User.ID+"> "+
			getRandomResponse(),
	)

	return err
}
