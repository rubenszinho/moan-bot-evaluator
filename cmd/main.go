package main

import (
	"log"
	"os"

	"discord-moan-bot/bot"
	"discord-moan-bot/config"
)

func main() {
	config.LoadEnv()

	token := os.Getenv("DISCORD_BOT_TOKEN")
	guildID := os.Getenv("GUILD_ID")

	if token == "" || guildID == "" {
		log.Fatal("Missing environment variables. Check .env file.")
	}

	bot.Start(token, guildID)
}
