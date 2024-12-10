package main

import (
	"fmt"
	"os"
	"strings"

	gladia "github.com/samber/go-gladia/pkg/client"
)

func main() {
	client := gladia.NewGladiaClient(os.Getenv("GLADIA_API_KEY"))

	transcription, err := client.GetTranscription(gladia.TranscriptionRequest{
		AudioURL:       "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		Diarization:    true,
		DetectLanguage: true,
	})
	if err != nil {
		panic(err)
	}

	// Print the transcription
	fmt.Printf("Languages: %s\n\n", strings.Join(transcription.Result.Transcription.Languages, ", "))
	fmt.Printf("Transcript: %s\n", transcription.Result.Transcription.FullTranscript)
}
