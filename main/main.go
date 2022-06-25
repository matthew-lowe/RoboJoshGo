package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/matthew-lowe/RoboJoshGo/commands"
	"github.com/matthew-lowe/RoboJoshGo/framework"

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

	session, err := discordgo.New("Bot " + token)
	defer session.Close()

	if err != nil {
		log.Fatal(err)
	}

	CmdRegistry = framework.NewCommandHandler()

	session.AddHandlerOnce(func(session *discordgo.Session, event *discordgo.Ready) {
		registerCommands(session)
	})

	session.AddHandler(interactionCreate)

	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bot now running. Press CTRL+C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func interactionCreate(session *discordgo.Session, event *discordgo.InteractionCreate) {
	user := event.User

	channel, err := session.Channel(event.ChannelID)
	if err != nil {
		return
	}

	guild, err := session.Guild(event.GuildID)
	if err != nil {
		return
	}

	context := framework.Context{
		Session:     session,
		User:        user,
		Channel:     channel,
		Guild:       guild,
		Interaction: event,
	}

	name := event.ApplicationCommandData().Name
	handler := commands.Handlers[name]
	err = handler(&context)

	if err != nil {
		log.Println(err)
	}
}

func registerCommands(session *discordgo.Session) {
	for _, v := range commands.Commands {
		CmdRegistry.Register(v, session)
	}
}
