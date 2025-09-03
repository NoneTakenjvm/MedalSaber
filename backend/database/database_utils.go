package database

import (
	"context"
	
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Fetch a score from the database
func GetScore(platform int, leaderboardId string, playerId string) (Score, error) {
	document, err := fetchDocument(Collections.Scores, bson.M{
		"platform":      platform,
		"playerId":      playerId,
		"leaderboardId": leaderboardId,
	})
	if err != nil {
		return Score{}, err
	}
	var score Score
	err = document.Decode(&score)
	if err != nil {
		return Score{}, err
	}
	return score, nil
}

// Fetch a player from the database
func GetPlayer(platform int, playerId string) (*Player, error) {
	document, err := fetchDocument(Collections.Players, bson.M{
		"platform": platform,
		"playerId": playerId,
	})
	if err != nil {
		return nil, err
	}
	var player Player
	err = document.Decode(&player)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

// Fetch a change from the database
func GetChanges(platform int, playerId string, page int) ([]Change, error) {
	cursor, err := fetchDocuments(Collections.Changes, bson.M{
		"platform": platform,
		"playerId": playerId,
	}, options.Find().SetSkip(int64(page*10)).SetLimit(10))
	if err != nil {
		return []Change{}, err
	}
	defer cursor.Close(context.Background())
	var changes []Change
	if err = cursor.All(context.Background(), &changes); err != nil {
		return []Change{}, err
	}
	return changes, nil
}

// Return whether the provided score is within the top 10 for that leaderboard
func IsWithinTopTen(platform int, leaderboardId string, country string, score int) (bool, error) {
	cursor, err := fetchDocuments(Collections.Scores, bson.M{
		"platform":      platform,
		"leaderboardId": leaderboardId,
		"country":       country,
	}, options.Find().SetSkip(int64(9)).SetLimit(1))
	if err != nil {
		return false, err
	}
	defer cursor.Close(context.Background())
	var changes []Score
	if err = cursor.All(context.Background(), &changes); err != nil {
		return false, err
	}
	// Check if we actually got any results
	if len(changes) == 0 {
		// No 10th place score exists, so any score is within top 10
		return true, nil
	}
	return score > changes[0].Score, nil
}

// Get the top 10 scores for a leaderboard
func GetTopTen(platform int, leaderboardId string, country string, playerId string) ([]Score, error) {
	cursor, err := fetchDocuments(Collections.Scores, bson.M{
		"platform":      platform,
		"leaderboardId": leaderboardId,
		"country":       country,
	}, options.Find().SetLimit(10))
	if err != nil {
		return []Score{}, err
	}
	defer cursor.Close(context.Background())
	var scores []Score
	if err = cursor.All(context.Background(), &scores); err != nil {
		return []Score{}, err
	}
	return scores, bson.ErrNilReader
}

func UpdateMedalCount(platform int, playerId string, positionNew int, positionOld int) {
	
}

