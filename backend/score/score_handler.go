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

// The medal value of each (indexed) position in the leaderboard 
var MedalValues = map[int]int{
	0:  10,
	1:  8,
	2:  6,
	3:  5,
	4:  4,
	5:  3,
	6:  2,
	7:  1,
	8:  1,
	9: 1,
	// We need to specify the 11th position score because scores pushed out of the top 10
	// will need to know how many medals position 11 is worth, which is 0!
	10: 0,
}

// Generic score interface for all platforms
type ScoreMessage interface {
	GetScoreId() string
	GetPlayerId() string
	GetPlayerName() string
	GetLeaderboardId() string
	GetLeaderboardName() string
	GetDifficulty() string
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

// Handle the provided score for the given country
//
// This function will:
// - award medals to the player who set the score
// - take medals from players who have been pushed out of the top 10
// - remove any score pushed from the top 10
// - update medal counts for all affected players
func handleForCountry(incomingScore ScoreMessage, country string) {
	// Get the country the score was set from, is it within top 10?
	isWithinTopTen, err := database.IsWithinTopTen(incomingScore.GetPlatform(), incomingScore.GetLeaderboardId(), country, incomingScore.GetScore())
	if err != nil {
		log.Printf("error when checking if a score is within top 10: %s\n", err)
		return
	}
	// If not within the top 10, we don't care
	if !isWithinTopTen {
		return
	}
	topTenScores, err := database.GetTopTen(incomingScore.GetPlatform(), incomingScore.GetLeaderboardId(), country, incomingScore.GetPlayerId())
	if err != nil {
		log.Printf("error when getting top 10 scores: %s\n", err)
		return
	}
	medalDeltas := make(map[string]int)
	position := getScorePositionInTopTen(topTenScores, incomingScore)
	alreadyPresent := isPlayerWithinTopTen(topTenScores, incomingScore.GetPlayerId())
	// The player has improved their score, but their position on the leaderboard hasn't changed
	// Or, the score is not in the top 10 at all
	if alreadyPresent == position || position == -1 {
		log.Printf("score from player %s (platform: %d, id: %s, country: %s) on leaderboard %s (difficulty: %s) was not improved or not within country top 10",
			incomingScore.GetPlayerName(), incomingScore.GetPlatform(), incomingScore.GetPlayerId(), country, incomingScore.GetLeaderboardName(), incomingScore.GetDifficulty())
		return
	}
	// Calculate medal deltas for all affected players
	calculateMedalDeltas(medalDeltas, topTenScores, incomingScore, position, alreadyPresent)
	// Handle score removal if we're at capacity and adding a new player
	if alreadyPresent == -1 && len(topTenScores) >= 10 {
		// Remove the 10th place score (index 9) since we're adding a new player
		removedScore := topTenScores[9]
		database.DeleteDocument(database.Collections.Scores, bson.M{
			"scoreId":  removedScore.ScoreId,
			"platform": removedScore.Platform,
		})
	}
	// Insert the new score into the database
	newScore := convertIntoDatabaseScore(country, incomingScore)
	err = database.InsertDocument(database.Collections.Scores, newScore)
	if err != nil {
		log.Printf("error when inserting new score: %s\n", err)
	}
	// Finally, handle the medal changes for all players
	handleMedalChanges(medalDeltas, incomingScore)
	log.Printf("the score from player %s (platform: %d, id: %s, country: %s) on leaderboard %s (difficulty: %s) has been handled! the player earned position %d",
		incomingScore.GetPlayerName(), incomingScore.GetPlatform(), incomingScore.GetPlayerId(), country, incomingScore.GetLeaderboardName(), incomingScore.GetDifficulty(), position)
}

// --- various single use helper functions to help organise code

// Return the position the player is in within the top 10 scores
// Will return -1 if the player is not already within the top 10
func isPlayerWithinTopTen(topTenScores []database.Score, playerId string) int {
	for i, score := range topTenScores {
		if score.PlayerId == playerId {
			return i
		}
	}
	return -1
}

// Return what position within the top 10 the incoming score would be
func getScorePositionInTopTen(topTenScores []database.Score, incomingScore ScoreMessage) int {
	if len(topTenScores) == 0 {
		return 0
	}
	for i, score := range topTenScores {
		if incomingScore.GetScore() > score.Score {
			return i
		}
	}
	return -1
}

// Convert the incoming score into a database score
func convertIntoDatabaseScore(country string, incomingScore ScoreMessage) database.Score {
	return database.Score{
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
}

// Calculate medal deltas for all affected players
func calculateMedalDeltas(medalDeltas map[string]int, topTenScores []database.Score, incomingScore ScoreMessage, position int, alreadyPresent int) {
	// Calculate how many medals each player moved down will lose
	for i := position + 1; i < len(topTenScores); i++ {
		// If the player being checked is the player who set the score, we need to remove their medals
		// entirely. They will be awarded new medals for their improved score.
		if topTenScores[i].PlayerId == incomingScore.GetPlayerId() {
			medalDeltas[topTenScores[i].PlayerId] = -MedalValues[alreadyPresent]
			continue
		}
		medalDeltas[topTenScores[i].PlayerId] = MedalValues[i-1] - MedalValues[i]
	}
	// Next, add the medals for the player who set the score
	medalDeltas[incomingScore.GetPlayerId()] += MedalValues[position]
}

// Handle medal changes for all players in the map
func handleMedalChanges(medalDeltas map[string]int, incomingScore ScoreMessage) {
	// Apply the medal deltas to all the players in the map
	for playerId, delta := range medalDeltas {
		player, err := database.GetPlayer(incomingScore.GetPlatform(), playerId)
		if err != nil {
			log.Printf("error when getting player: %s\n", err)
			continue
		}
		player.Medals += delta
		if err = database.UpdateDocument(
			database.Collections.Players,
			bson.M{"playerId": playerId, "platform": incomingScore.GetPlatform()},
			bson.M{"$set": bson.M{"medals": player.Medals}}); err != nil {
			log.Printf("error when updating player: %s\n", err)
		}
	}
}
