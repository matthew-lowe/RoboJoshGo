package framework

import (
	"github.com/bwmarrin/discordgo"
	"sort"
)

type (
	History struct {
		deletions map[string]*discordgo.Message
	}
)

func (history *History) AddMessage(message *discordgo.Message) {
	history.deletions[message.ID] = message
}

// Gets the latest history
func (history *History) GetMessage(id string) *discordgo.Message {
	return history.deletions[id]
}

func (history *History) GetLastMessage(index int) *discordgo.Message {
	keys := make([]string, len(history.deletions))

	i := 0
	for k := range history.deletions {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	return history.deletions[keys[len(keys)-index-1]]
}

func NewHistory() *History {
	return &History{make(map[string]*discordgo.Message)}
}
