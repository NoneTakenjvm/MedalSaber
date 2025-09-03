package score

import (
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"nonetaken.dev/medalsaber/database"
)

// Create constant values for each platform
const (
	ScoresaberPlatform int = 1
	BeatleaderPlatform int = 2
)

// The medal value of each position in the leaderboard
var MedalValues = map[int]int{
		1: 10,
		2: 8,
		3: 6,
		4: 5,
		5: 4,
		6: 3,
		7: 2,
		8: 1,
		9: 1,
		10: 1,
}

// Generic score interface for all platforms
type ScoreMessage interface {
	GetScoreId() string
	GetPlayerId() string
	GetLeaderboardId() string
	GetCountry() string
	GetScore() int
	GetMaxScore() int
	GetPlatform() int
	IsRanked() bool
	GetTimestamp() int64
	GetModifiers() string
	GetBadCuts() int
	GetMissedNotes() int
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
	// If the score is not ranked, we don't care
	if !incomingScore.IsRanked() {
		return
	}
	// Handle for the country the score was set from and for the world
	handleForCountry(incomingScore, incomingScore.GetCountry())
	handleForCountry(incomingScore, "Global")
}

func handleForCountry(incomingScore ScoreMessage, country string) {
	// Get the country the score was set from, is it within top 10?
	isWithinTopTen, err := database.IsWithinTopTen(incomingScore.GetPlatform(), incomingScore.GetLeaderboardId(), country, incomingScore.GetScore())
	if err != nil {
		log.Printf("error when checking if a score is within top 10: %s\n", err)
	}
	// If not within the top 10, we don't care
	if !isWithinTopTen {
		return
	}
	topTenScores, err := database.GetTopTen(incomingScore.GetPlatform(), incomingScore.GetLeaderboardId(), country, incomingScore.GetPlayerId())
	if err != nil {
		log.Printf("error when getting top 10 scores: %s\n", err)
	}
	// Find where this score belongs within the top 10
	var position int
	for _, score := range topTenScores {
		if incomingScore.GetScore() > score.Score{
			break
		}
		position++
	}
	// Store the score at position 10 (if the top 10 is full) so it can be removed later
	var removedScore *database.Score = nil
	if len(topTenScores) == 10 {
		removedScore = &topTenScores[9]
		// Delete the removed score here, we will never need it again
		database.DeleteDocument(database.Collections.Scores, bson.M{
			"scoreId": removedScore.ScoreId,
			"platform": incomingScore.GetPlatform(),
		})
	}
	// If the score will be inserted at index 0-8, we need to move scores around to freeup that position in the array
	if 8 >= position {
		topTenScores = append(topTenScores[:position+1], topTenScores[position:]...)
		topTenScores = topTenScores[:10] // Trim the array to keep it at 10 scores
	}
	// Insert the new score to the array
	topTenScores[position] = database.Score{
		ScoreId:       incomingScore.GetScoreId(),
		PlayerId:      incomingScore.GetPlayerId(),
		Country:       country,
		LeaderboardId: incomingScore.GetLeaderboardId(),
		Platform:      incomingScore.GetPlatform(),
		Score:         incomingScore.GetScore(),
		MaxScore:      incomingScore.GetMaxScore(),
		Timestamp:     incomingScore.GetTimestamp(),
		Modifiers:     incomingScore.GetModifiers(),
		BadCuts:       incomingScore.GetBadCuts(),
		MissedNotes:   incomingScore.GetMissedNotes(),
	}

	// Update the medal counts for players who were shifted down
	if len(topTenScores) > position {

	}

	// update each affected spot's medal acount as well as the player who set the score
	// update the top 10 document
	// update the player documents
	// done
}
