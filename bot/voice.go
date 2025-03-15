package bot

import (
	"fmt"

	"discord-moan-bot/audio"

	"github.com/bwmarrin/discordgo"
)

func voiceCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channelID := "TARGET_VOICE_CHANNEL_ID"

	vc, err := s.ChannelVoiceJoin(i.GuildID, channelID, false, false)
	if err != nil {
		fmt.Println("Error joining voice channel:", err)
		return
	}
	defer vc.Disconnect()

	audio.PlayAudio(vc, "bip.mp3")
	audio.RecordAudio(vc, "recorded.wav")

	result := audio.Evaluate("recorded.wav")
	s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("Evaluation: %s", result))
}
