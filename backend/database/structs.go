package database

type Score struct {
	ScoreId       string `bson:"scoreId"`
	PlayerId      string `bson:"playerId"`
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
func (score *Score) GetPlayer(region string) *Player {
	player, err := GetPlayer(score.Platform, score.PlayerId, region, "", true)
	if err != nil {
		return nil
	}
	return player
}

// Player struct ----------------

type Player struct {
	PlayerId string `bson:"playerId"`
	Platform int    `bson:"platform"`
	Region   string `bson:"region"`
	Medals   int    `bson:"medals"`
	Username string `bson:"username"`
}

// Change struct ----------------

type Change struct {
	Platform                 int    `bson:"platform"`
	PlayerId                 string `bson:"playerId"`
	Region                   string `bson:"region"`
	Timestamp                int64  `bson:"timestamp"`
	MedalChange              int    `bson:"medalChange"`
	ResponsibleLeaderboardId string `bson:"responsibleLeaderboardId"`
	ResponsiblePlayerId      string `bson:"responsiblePlayerId"`
	ResponsibleScoreId       string `bson:"responsibleScoreId"`
}
