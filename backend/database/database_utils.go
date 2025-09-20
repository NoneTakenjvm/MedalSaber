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
	document, err := fetchDocument(Collections.Scores, bson.M{
		"platform":      platform,
		"scoreId":       scoreId,
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

// Fetch a player from the database, creating one if they don't exist
func GetPlayer(platform int, country string, playerId string, createIfAbsent bool) (*Player, error) {
	document, err := fetchDocument(Collections.Players, bson.M{
		"platform": platform,
		"playerId": playerId,
	})
	// Create the player document if they don't exist already
	if err != nil {
		// Check if it's a "not found" error specifically
		if err == mongo.ErrNoDocuments {
			// Use mutex to prevent race conditions when creating players
			playerCreationMutex.Lock()
			defer playerCreationMutex.Unlock()

			// Double-check: try to fetch again in case another goroutine created it
			document, err = fetchDocument(Collections.Players, bson.M{
				"platform": platform,
				"playerId": playerId,
			})
			if err == nil {
				// Player was created by another goroutine, decode and return
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
					Country:  country,
					Medals:   0,
				}
				err = InsertDocument(Collections.Players, newPlayer)
				if err != nil {
					return nil, err
				}
				return &newPlayer, nil
			}
		}
		// Some other error occurred, return it
		return nil, err
	}
	var player Player
	err = document.Decode(&player)
	if err != nil {
		return nil, err
	}
	// Create a defer to check if the player has changed country
	defer func() {
		if player.Country != country {
			// Update the player's country
			UpdateDocument(Collections.Players, bson.M{"playerId": playerId, "platform": platform}, bson.M{"$set": bson.M{"country": country}})
			// Update the player's scores to the new country
			UpdateManyDocuments(Collections.Scores, bson.M{"playerId": playerId, "platform": platform}, bson.M{"$set": bson.M{"country": country}})
		}
	}()
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
	return scores, nil
}
