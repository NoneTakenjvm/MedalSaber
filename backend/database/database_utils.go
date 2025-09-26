package database

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Lock for player creation to prevent race conditions
var playerCreationMutex sync.Mutex

// Fetch a score from the database
func GetScore(platform int, scoreId string) (Score, error) {
	document, err := FetchDocument(Collections.Scores, bson.M{
		"platform": platform,
		"scoreId":  scoreId,
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

// Fetch all of a player's scores from the database
func GetPlayerScores(player *Player, page int, before int64, after int64) ([]Score, error) {
	// Build the mongo filter
	filter := bson.M{
		"platform": player.Platform,
		"playerId": player.PlayerId,
		"region":   player.Region,
	}
	// Add the before and after filters
	if before != 0 || after != 0 {
		timestamp := bson.M{}
		if before != 0 {
			timestamp["$lte"] = before
		}
		if after != 0 {
			timestamp["$gte"] = after
		}
		filter["timestamp"] = timestamp
	}
	// Fetch the player's scores from the database
	cursor, err := FetchDocuments(Collections.Scores, filter, options.Find().SetSkip(int64(page*10)).SetLimit(10))
	if err != nil {
		return []Score{}, err
	}
	defer cursor.Close(context.Background())
	var scores []Score
	if err = cursor.All(context.Background(), &scores); err != nil {
		return []Score{}, err
	}
	return scores, nil
}

// Fetch a change from the database
func GetChanges(platform int, region string, playerId string, page int, before int64, after int64) ([]Change, error) {
	// Build the mongo filter
	filter := bson.M{
		"platform": platform,
		"playerId": playerId,
		"region":   region,
	}
	// Add the before and after filters
	if before != 0 || after != 0 {
		timestamp := bson.M{}
		if before != 0 {
			timestamp["$lte"] = before
		}
		if after != 0 {
			timestamp["$gte"] = after
		}
		filter["timestamp"] = timestamp
	}
	// Fetch the changes from the database
	cursor, err := FetchDocuments(Collections.Changes, filter, options.Find().SetSkip(int64(page*10)).SetLimit(10))
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
func IsWithinTopTen(platform int, leaderboardId string, region string, score int) (bool, error) {
	cursor, err := FetchDocuments(Collections.Scores, bson.M{
		"platform":      platform,
		"leaderboardId": leaderboardId,
		"region":        region,
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
func GetTopTenScores(platform int, region string, leaderboardId string, page int64) ([]Score, error) {
	filter := bson.M{
		"platform":      platform,
		"region":        region,
		"leaderboardId": leaderboardId,
	}
	cursor, err := FetchDocuments(Collections.Scores, filter, options.Find().SetLimit(page*10))
	if err != nil {
		return []Score{}, err
	}
	defer cursor.Close(context.Background())
	var scores []Score
	if err = cursor.All(context.Background(), &scores); err != nil {
		return []Score{}, err
	}
	return scores, nil
}

// Get the top 10 medal holders for a region
func GetTopTenMedalHolders(platform int, region string, page int64) ([]Player, error) {
	filter := bson.M{
		"platform": platform,
		"region":   region,
	}
	cursor, err := FetchDocuments(Collections.Players, filter, options.Find().SetSkip(int64(page*10)).SetLimit(10))
	if err != nil {
		return []Player{}, err
	}
	defer cursor.Close(context.Background())
	var players []Player
	if err = cursor.All(context.Background(), &players); err != nil {
		return []Player{}, err
	}
	return players, nil
}

// Fetch a player from the database, optionally creating one if they don't exist
func GetPlayer(platform int, region string, playerId string, username string, createIfAbsent bool) (*Player, error) {
	document, err := FetchDocument(Collections.Players, bson.M{
		"platform": platform,
		"playerId": playerId,
		"region":   region,
	})
	// Create the player document if they don't exist already
	if err == mongo.ErrNoDocuments {
		// Use a mutex to prevent race conditions when creating players
		playerCreationMutex.Lock()
		defer playerCreationMutex.Unlock()
		// Try to fetch again in case another goroutine created it
		document, err = FetchDocument(Collections.Players, bson.M{
			"platform": platform,
			"playerId": playerId,
			"region":   region,
		})
		// The player was created by another goroutine, decode and return
		if err == nil {
			var player Player
			err = document.Decode(&player)
			if err != nil {
				return nil, err
			}
			return &player, nil
		}
		// Still not found, create a new one
		if err == mongo.ErrNoDocuments && createIfAbsent {
			newPlayer := Player{
				PlayerId: playerId,
				Platform: platform,
				Region:   region,
				Medals:   0,
				Username: username,
			}
			err = InsertDocument(Collections.Players, newPlayer)
			if err != nil {
				return nil, err
			}
			return &newPlayer, nil
		}
		return nil, err
	}
	// Some other error occured, return it
	if err != nil {
		// Some other error occurred, return it
		return nil, err
	}
	// Parse the found or created player
	var player Player
	err = document.Decode(&player)
	if err != nil {
		return nil, err
	}
	return &player, nil
}
