package audio

import (
	"bufio"
	"fmt"
	"os/exec"
	"time"

	"github.com/bwmarrin/discordgo"
)

func PlayAudio(vc *discordgo.VoiceConnection, file string) {
	cmd := exec.Command("ffmpeg", "-i", file, "-f", "s16le", "-ar", "48000", "-ac", "2", "pipe:1")
	stdout, _ := cmd.StdoutPipe()

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting ffmpeg:", err)
		return
	}

	opusEncoder := NewOpusEncoder(vc.OpusSend)
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		opusEncoder.Encode(scanner.Bytes())
	}

	cmd.Wait()
}

func RecordAudio(vc *discordgo.VoiceConnection, outputFile string) {
	cmd := exec.Command("ffmpeg", "-f", "pulse", "-i", "default", "-t", "10", "-ac", "2", "-ar", "48000", outputFile)

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting ffmpeg recording:", err)
		return
	}

	time.Sleep(10 * time.Second)
	cmd.Wait()

	fmt.Println("Audio recorded:", outputFile)
}
