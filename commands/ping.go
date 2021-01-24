package commands

import (
	"github.com/matthewlowe/Robojosh/framework"
)

func PingCommand(context *framework.Context) error {
	_, err := context.Session.ChannelMessageSend(context.TextChannel.ID, "Pong!")

	return err
}
