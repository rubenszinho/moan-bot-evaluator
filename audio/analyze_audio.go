package audio

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Evaluate extracts audio features and evaluates the moan quality
func Evaluate(audioFile string) string {
	// Define the OpenSMILE config file for pitch & intensity analysis
	configFile := "config/emobase.conf" // Use an OpenSMILE config suited for prosody analysis

	// Run OpenSMILE command to extract audio features
	cmd := exec.Command("SMILExtract", "-C", configFile, "-I", audioFile, "-O", "features.csv")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running OpenSMILE:", err)
		return "Error analyzing audio"
	}

	// Parse OpenSMILE output
	features := parseFeatures("features.csv")

	// Evaluate the moan based on extracted audio features
	evaluation := evaluateMoan(features)
	return evaluation
}

// parseFeatures reads OpenSMILE output and extracts pitch, duration, intensity
func parseFeatures(csvFile string) map[string]float64 {
	file, err := os.ReadFile(csvFile)
	if err != nil {
		fmt.Println("Error reading features file:", err)
		return nil
	}

	// Extract pitch, intensity, and duration from CSV
	lines := strings.Split(string(file), "\n")
	features := make(map[string]float64)

	for _, line := range lines {
		if strings.Contains(line, "F0final") { // Pitch
			fields := strings.Split(line, ",")
			features["pitch"], _ = strconv.ParseFloat(fields[1], 64)
		}
		if strings.Contains(line, "pcm_intensity") { // Intensity
			fields := strings.Split(line, ",")
			features["intensity"], _ = strconv.ParseFloat(fields[1], 64)
		}
	}

	// Estimate duration (mocked as OpenSMILE does not directly return it)
	features["duration"] = 10.0 // Assume full 10s recording

	return features
}

// evaluateMoan provides a qualitative score based on extracted features
func evaluateMoan(features map[string]float64) string {
	if features == nil {
		return "Error processing features"
	}

	pitch := features["pitch"]
	intensity := features["intensity"]
	duration := features["duration"]

	// Example scoring criteria
	if pitch > 200 && intensity > 60 && duration >= 5 {
		return "🔥 Passionate moan detected! High pitch and intensity."
	} else if pitch > 150 && intensity > 40 {
		return "🙂 Moderate moan. Could be more expressive."
	} else {
		return "😐 Weak moan. Try adding more emotion!"
	}
}
