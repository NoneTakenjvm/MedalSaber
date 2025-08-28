package score

/*
 * Courtesy of https://git.fascinated.cc/Fascinated/yet-another-bs-tracker/src/branch/master/src/leaderboards/scoresaber/scoresaber.go
 */

/*
 * Structs for the ScoreSaber API
 */
type Player struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	ProfilePicture    string     `json:"profilePicture"`
	Bio               string     `json:"bio"`
	Country           string     `json:"country"`
	PerformancePoints float64    `json:"pp"`
	Rank              int        `json:"rank"`
	CountryRank       int        `json:"countryRank"`
	Role              int        `json:"role"`
	Badges            []Badge    `json:"badges"`
	Histories         string     `json:"histories"`
	Permissions       int        `json:"permissions"`
	Banned            bool       `json:"banned"`
	Inactive          bool       `json:"inactive"`
	ScoreStats        ScoreStats `json:"scoreStats"`
	FirstSeen         string     `json:"firstSeen"`
}

type Badge struct {
	Image       string `json:"image"`
	Description string `json:"description"`
}

type ScoreStats struct {
	TotalScore            int     `json:"totalScore"`
	TotalRankedScore      int     `json:"totalRankedScore"`
	AverageRankedAccuracy float64 `json:"averageRankedAccuracy"`
	TotalPlayCount        int     `json:"totalPlayCount"`
	RankedPlayCount       int     `json:"rankedPlayCount"`
	ReplaysWatched        int     `json:"replaysWatched"`
}

/*
 * Structs for the ScoreSaber websocket
 */
type IncomingMessage struct {
	CommandName string `json:"commandName"`
}

type IncomingMessageWithScore struct {
	CommandName string        `json:"commandName"`
	Score       IncomingScore `json:"commandData"`
}

type IncomingScore struct {
	Score       Score       `json:"score"`
	Leaderboard Leaderboard `json:"leaderboard"`
}

type Score struct {
	ID                    int                   `json:"id"`
	LeaderboardPlayerInfo LeaderboardPlayerInfo `json:"leaderboardPlayerInfo"`
	Rank                  int                   `json:"rank"`
	BaseScore             float64               `json:"baseScore"`
	ModifiedScore         int                   `json:"modifiedScore"`
	PP                    float64               `json:"pp"`
	Weight                float64               `json:"weight"`
	Modifiers             string                `json:"modifiers"`
	Multiplier            float64               `json:"multiplier"`
	BadCuts               int                   `json:"badCuts"`
	MissedNotes           int                   `json:"missedNotes"`
	MaxCombo              int                   `json:"maxCombo"`
	FullCombo             bool                  `json:"fullCombo"`
	HMD                   int                   `json:"hmd"`
	TimeSet               string                `json:"timeSet"`
	HasReplay             bool                  `json:"hasReplay"`
	DeviceHMD             string                `json:"deviceHmd"`
	DeviceControllerLeft  string                `json:"deviceControllerLeft"`
	DeviceControllerRight string                `json:"deviceControllerRight"`
}

type Leaderboard struct {
	ID                int        `json:"id"`
	SongHash          string     `json:"songHash"`
	SongName          string     `json:"songName"`
	SongSubName       string     `json:"songSubName"`
	SongAuthorName    string     `json:"songAuthorName"`
	LevelAuthorName   string     `json:"levelAuthorName"`
	Difficulty        Difficulty `json:"difficulty"`
	MaxScore          float64    `json:"maxScore"`
	CreatedDate       string     `json:"createdDate"`
	RankedDate        string     `json:"rankedDate"`
	QualifiedDate     string     `json:"qualifiedDate"`
	LovedDate         string     `json:"lovedDate"`
	Ranked            bool       `json:"ranked"`
	Qualified         bool       `json:"qualified"`
	Loved             bool       `json:"loved"`
	MaxPP             float64    `json:"maxPP"`
	Stars             float64    `json:"stars"`
	Plays             int        `json:"plays"`
	DailyPlays        int        `json:"dailyPlays"`
	PositiveModifiers bool       `json:"positiveModifiers"`
	PlayerScore       string     `json:"playerScore"`
	CoverImage        string     `json:"coverImage"`
	// Difficulties      string     `json:"difficulties"`
}

type LeaderboardPlayerInfo struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profilePicture"`
	Country        string `json:"country"`
	Permissions    int    `json:"permissions"`
	Badges         string `json:"badges"`
	Role           string `json:"role"`
}

type Difficulty struct {
	LeaderboardID int    `json:"leaderboardId"`
	Difficulty    int    `json:"difficulty"`
	GameMode      string `json:"gameMode"`
	DifficultyRaw string `json:"difficultyRaw"`
}
