package score

import (
	"encoding/json"
	"log"
)

// Create constant values for each platform
const (
	ScoresaberPlatform int = 1
	BeatleaderPlatform int = 2
)

// Generic score interface for all platforms
type ScoreMessage interface {
	GetPlayerId() string
	GetLeaderboardId() string
	GetCountry() string
	GetScore() int
	GetPlatform() int
}

func HandleScore(platform int, message []byte) {
	var incomingScore ScoreMessage
	// Handle scores from ScoreSaber
	if platform == ScoresaberPlatform {
		var scoresaberMessage IncomingMessageWithScore
		err := json.Unmarshal(message, &scoresaberMessage)
		if err != nil {
			log.Printf("Error while parsing Scoresaber message: %s\n", err)
		}
		incomingScore = &scoresaberMessage
	}
	// Handle scores from BeatLeader
	if platform == BeatleaderPlatform {
		var beatleaderMessage BeatLeaderResponse
		err := json.Unmarshal(message, &beatleaderMessage)
		if err != nil {
			log.Printf("Error while parsing Beatleader message: %s\n", err)
		}
		incomingScore = &beatleaderMessage
	}
	if incomingScore.GetCountry() == "lol" {

	}
}
