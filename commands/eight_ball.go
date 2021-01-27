package commands

import (
	"github.com/matthewlowe/RoboJoshGo/framework"
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
	_, err := context.Reply(getRandomResponse())

	return err
}
