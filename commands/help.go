package commands

import (
	"github.com/matthew-lowe/RoboJoshGo/framework"
)

func HelpCommand(context *framework.Context) error {
	err := context.Reply("Fuck you")
	return err
}
