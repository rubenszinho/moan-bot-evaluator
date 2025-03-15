package audio

import (
	"encoding/binary"
	"fmt"
	"time"

	"layeh.com/gopus"
)

// OpusEncoder handles encoding PCM data into Opus frames for Discord
type OpusEncoder struct {
	sendChannel chan []byte
	encoder     *gopus.Encoder
}

// NewOpusEncoder initializes a new Opus encoder
func NewOpusEncoder(sendChan chan []byte) *OpusEncoder {
	encoder, err := gopus.NewEncoder(48000, 2, gopus.Audio)
	if err != nil {
		fmt.Println("Error initializing Opus encoder:", err)
		return nil
	}
	encoder.SetBitrate(64000)

	return &OpusEncoder{
		sendChannel: sendChan,
		encoder:     encoder,
	}
}

func (o *OpusEncoder) Encode(pcmData []byte) {
	if len(pcmData)%2 != 0 {
		fmt.Println("Invalid PCM data length")
		return
	}

	// Convert []byte to []int16
	pcmInt16 := make([]int16, len(pcmData)/2)
	for i := 0; i < len(pcmData); i += 2 {
		pcmInt16[i/2] = int16(binary.LittleEndian.Uint16(pcmData[i : i+2]))
	}

	// Encode to Opus
	opusData, err := o.encoder.Encode(pcmInt16, 960, 2) // 960 samples per frame, 2 channels
	if err != nil {
		fmt.Println("Error encoding to Opus:", err)
		return
	}

	o.sendChannel <- opusData
	time.Sleep(20 * time.Millisecond) // Discord requires Opus packets every 20ms
}
