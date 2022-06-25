package commands

import (
	"github.com/matthew-lowe/RoboJoshGo/framework"
)

func PingCommand(context *framework.Context) error {
	err := context.Reply("Pong!")

	return err
}
