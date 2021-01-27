package commands

import (
	"github.com/matthewlowe/RoboJoshGo/framework"
)

func PingCommand(context *framework.Context) error {
	_, err := context.Reply("Pong!")

	return err
}
