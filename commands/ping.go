package commands

import (
	"github.com/matthewlowe/Robojosh/framework"
)

func PingCommand(context *framework.Context) {
	context.Discord.ChannelMessageSend(context.TextChannel.ID, "Pong!")
}
