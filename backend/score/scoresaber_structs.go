package score

import (
	"log"
	"strconv"
)

/*
 * Structs for the ScoreSaber API
 */
type IncomingMessageWithScore struct {
	CommandName string                  `json:"commandName"`
	Score       ScoresaberIncomingScore `json:"commandData"`
}

// Implement functions for the IncomingMessageWithScore struct so it can become a ScoreMessage interface
func (message *IncomingMessageWithScore) GetPlayerId() string {
	return message.Score.Score.LeaderboardPlayerInfo.ID
}
func (message *IncomingMessageWithScore) GetLeaderboardId() string {
	return strconv.Itoa(message.Score.Leaderboard.ID)
}
func (message *IncomingMessageWithScore) GetCountry() string {
	return message.Score.Score.LeaderboardPlayerInfo.Country
}
func (message *IncomingMessageWithScore) GetScore() int {
	return message.Score.Score.ModifiedScore
}
func (message *IncomingMessageWithScore) GetPlatform() int {
	return ScoresaberPlatform
}
func (message *IncomingMessageWithScore) IsRanked() bool {
	return message.Score.Score.PP > 0
}
func (message *IncomingMessageWithScore) GetScoreId() string {
	return strconv.Itoa(message.Score.Score.ID)
}
func (message *IncomingMessageWithScore) GetMaxScore() int {
	return int(message.Score.Leaderboard.MaxScore)
}
func (message *IncomingMessageWithScore) GetTimestamp() int64 {
	num, err := strconv.ParseInt(message.Score.Score.TimeSet, 10, 64)
	if err != nil {
		log.Printf("error when parsing timestamp: %s\n", err)
	}
	return num
}
func (message *IncomingMessageWithScore) GetModifiers() string {
	return message.Score.Score.Modifiers
}
func (message *IncomingMessageWithScore) GetBadCuts() int {
	return message.Score.Score.BadCuts
}
func (message *IncomingMessageWithScore) GetMissedNotes() int {
	return message.Score.Score.MissedNotes
}

type ScoresaberPlayer struct {
	ID                string               `json:"id"`
	Name              string               `json:"name"`
	ProfilePicture    string               `json:"profilePicture"`
	Bio               string               `json:"bio"`
	Country           string               `json:"country"`
	PerformancePoints float64              `json:"pp"`
	Rank              int                  `json:"rank"`
	CountryRank       int                  `json:"countryRank"`
	Role              int                  `json:"role"`
	Badges            []ScoresaberBadge    `json:"badges"`
	Histories         string               `json:"histories"`
	Permissions       int                  `json:"permissions"`
	Banned            bool                 `json:"banned"`
	Inactive          bool                 `json:"inactive"`
	ScoreStats        ScoresaberScoreStats `json:"scoreStats"`
	FirstSeen         string               `json:"firstSeen"`
}

type ScoresaberBadge struct {
	Image       string `json:"image"`
	Description string `json:"description"`
}

type ScoresaberScoreStats struct {
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
type ScoresaberIncomingMessage struct {
	CommandName string `json:"commandName"`
}

type ScoresaberIncomingScore struct {
	Score       ScoresaberScore       `json:"score"`
	Leaderboard ScoresaberLeaderboard `json:"leaderboard"`
}

type ScoresaberScore struct {
	ID                    int                             `json:"id"`
	LeaderboardPlayerInfo ScoresaberLeaderboardPlayerInfo `json:"leaderboardPlayerInfo"`
	Rank                  int                             `json:"rank"`
	BaseScore             float64                         `json:"baseScore"`
	ModifiedScore         int                             `json:"modifiedScore"`
	PP                    float64                         `json:"pp"`
	Weight                float64                         `json:"weight"`
	Modifiers             string                          `json:"modifiers"`
	Multiplier            float64                         `json:"multiplier"`
	BadCuts               int                             `json:"badCuts"`
	MissedNotes           int                             `json:"missedNotes"`
	MaxCombo              int                             `json:"maxCombo"`
	FullCombo             bool                            `json:"fullCombo"`
	HMD                   int                             `json:"hmd"`
	TimeSet               string                          `json:"timeSet"`
	HasReplay             bool                            `json:"hasReplay"`
	DeviceHMD             string                          `json:"deviceHmd"`
	DeviceControllerLeft  string                          `json:"deviceControllerLeft"`
	DeviceControllerRight string                          `json:"deviceControllerRight"`
}

type ScoresaberLeaderboard struct {
	ID                int                  `json:"id"`
	SongHash          string               `json:"songHash"`
	SongName          string               `json:"songName"`
	SongSubName       string               `json:"songSubName"`
	SongAuthorName    string               `json:"songAuthorName"`
	LevelAuthorName   string               `json:"levelAuthorName"`
	Difficulty        ScoresaberDifficulty `json:"difficulty"`
	MaxScore          float64              `json:"maxScore"`
	CreatedDate       string               `json:"createdDate"`
	RankedDate        string               `json:"rankedDate"`
	QualifiedDate     string               `json:"qualifiedDate"`
	LovedDate         string               `json:"lovedDate"`
	Ranked            bool                 `json:"ranked"`
	Qualified         bool                 `json:"qualified"`
	Loved             bool                 `json:"loved"`
	MaxPP             float64              `json:"maxPP"`
	Stars             float64              `json:"stars"`
	Plays             int                  `json:"plays"`
	DailyPlays        int                  `json:"dailyPlays"`
	PositiveModifiers bool                 `json:"positiveModifiers"`
	PlayerScore       string               `json:"playerScore"`
	CoverImage        string               `json:"coverImage"`
	// Difficulties      string     `json:"difficulties"`
}

type ScoresaberLeaderboardPlayerInfo struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profilePicture"`
	Country        string `json:"country"`
	Permissions    int    `json:"permissions"`
	Badges         string `json:"badges"`
	Role           string `json:"role"`
}

type ScoresaberDifficulty struct {
	LeaderboardID int    `json:"leaderboardId"`
	Difficulty    int    `json:"difficulty"`
	GameMode      string `json:"gameMode"`
	DifficultyRaw string `json:"difficultyRaw"`
}
