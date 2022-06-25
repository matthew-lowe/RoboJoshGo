package commands

import (
	"math/rand"

	"github.com/matthew-lowe/RoboJoshGo/framework"
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
	err := context.Reply(getRandomResponse())

	return err
}
