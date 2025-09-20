package database

type Score struct {
	ScoreId       string `bson:"scoreId"`
	PlayerId      string `bson:"playerId"`
	Country       string `bson:"country"`
	LeaderboardId string `bson:"leaderboardId"`
	Platform      int    `bson:"platform"`
	Score         int    `bson:"score"`
	MaxScore      int    `bson:"maxScore"`
	Timestamp     int64  `bson:"timestamp"`
	Modifiers     string `bson:"modifiers"`
	BadCuts       int    `bson:"badCuts"`
	MissedNotes   int    `bson:"missedNotes"`
}

// Get the player who set the score
func (score *Score) GetPlayer() *Player {
	player, err := GetPlayer(score.Platform, score.PlayerId, score.Country)
	if err != nil {
		return nil
	}
	return player
}

// Player struct ----------------

type Player struct {
	PlayerId string `bson:"playerId"`
	Platform int    `bson:"platform"`
	Country string  `bson:"country"`
	Medals   int    `bson:"medals"`
}

// Change struct ----------------

type Change struct {
	Platform                 int    `bson:"platform"`
	PlayerId                 string `bson:"playerId"`
	Timestamp                int64  `bson:"timestamp"`
	MedalChange              int    `bson:"medalChange"`
	ResponsibleLeaderboardId string `bson:"responsibleLeaderboardId"`
	ResponsiblePlayerId      string `bson:"responsiblePlayerId"`
	ResponsibleScoreId       string `bson:"responsibleScoreId"`
}
