package main

import (
	"fmt"
	"github.com/matthewlowe/Robojosh/commands"
	"github.com/matthewlowe/Robojosh/framework"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	CmdRegistry *framework.CommandRegistry
)

const (
	Prefix string = "%"
)

func loadToken() (string, error) {
	file, err := os.Open("token.txt")

	if err != nil {
		return "", err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(file)
	return string(b), err
}

func main() {
	token, err := loadToken()

	if err != nil {
		log.Fatal(err)
	}

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	CmdRegistry = framework.NewCommandHandler()
	registerCommands()

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bot now running. Press CTRL+C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = dg.Close() // We literally do not care if there's an error here, the program is closing anyway
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	switch {
	case message.Author.ID == session.State.User.ID:
		return
	case len(message.Content) < 2:
		return
	case message.Content[:len(Prefix)] != Prefix:
		return
	}

	args := strings.Fields(message.Content[len(Prefix):])
	name := args[0]

	user := message.Author

	channel, err := session.Channel(message.ChannelID)
	if err != nil {
		return
	}

	guild, err := session.Guild(message.GuildID)
	if err != nil {
		return
	}

	context := framework.Context{
		Session:     session,
		User:        user,
		TextChannel: channel,
		Guild:       guild,
		Message:     message.Message,
		Args:        args,
		CmdRegistry: CmdRegistry,
		Prefix:      Prefix,
	}

	command, found := CmdRegistry.Get(name)

	if !found {
		return
	}

	c := *command
	err = c(&context)

	if err != nil {
		log.Println(err)
	}
}

func registerCommands() {
	CmdRegistry.Register("ping", "Test if the bot is working", commands.PingCommand)
	CmdRegistry.Register("help", "View a list of commands", commands.HelpCommand)
	CmdRegistry.Register("color", "Generate a solid image color", commands.ColorCommand)
	CmdRegistry.Register("8ball", "Magic 8 ball!", commands.EightBallCommand)
}
