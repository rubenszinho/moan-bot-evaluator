package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var bot *discordgo.Session

func Start(token, guildID string) {
	var err error
	bot, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
	}

	bot.AddHandler(voiceCommandHandler)
	bot.Identify.Intents = discordgo.IntentsGuildVoiceStates

	err = bot.Open()
	if err != nil {
		log.Fatal("Error opening Discord session:", err)
	}

	fmt.Println("Bot is running. Press CTRL+C to exit.")
	select {}
}
