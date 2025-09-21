package score

import (
	"strconv"
)

type BeatLeaderResponse struct {
	ContextExtensions []ContextExtension    `json:"contextExtensions"`
	MyScore           any                   `json:"myScore"`
	ValidContexts     int                   `json:"validContexts"`
	Experience        int                   `json:"experience"`
	Leaderboard       BeatLeaderLeaderboard `json:"leaderboard"`
	AccLeft           float64               `json:"accLeft"`
	AccRight          float64               `json:"accRight"`
	ID                int                   `json:"id"`
	BaseScore         int                   `json:"baseScore"`
	ModifiedScore     int                   `json:"modifiedScore"`
	Accuracy          float64               `json:"accuracy"`
	PlayerID          string                `json:"playerId"`
	PP                float64               `json:"pp"`
	BonusPP           float64               `json:"bonusPp"`
	PassPP            float64               `json:"passPP"`
	AccPP             float64               `json:"accPP"`
	TechPP            float64               `json:"techPP"`
	Rank              int                   `json:"rank"`
	ResponseRank      int                   `json:"responseRank"`
	Country            string               `json:"country"`
	FcAccuracy        float64               `json:"fcAccuracy"`
	FcPp              float64               `json:"fcPp"`
	Weight            float64               `json:"weight"`
	Replay            string                `json:"replay"`
	Modifiers         string                `json:"modifiers"`
	BadCuts           int                   `json:"badCuts"`
	MissedNotes       int                   `json:"missedNotes"`
	BombCuts          int                   `json:"bombCuts"`
	WallsHit          int                   `json:"wallsHit"`
	Pauses            int                   `json:"pauses"`
	FullCombo         bool                  `json:"fullCombo"`
	Platform          string                `json:"platform"`
	MaxCombo          int                   `json:"maxCombo"`
	MaxStreak         int                   `json:"maxStreak"`
	HMD               int                   `json:"hmd"`
	Controller        int                   `json:"controller"`
	LeaderboardID     string                `json:"leaderboardId"`
	Timeset           string                `json:"timeset"`
	Timepost          int64                 `json:"timepost"`
	ReplaysWatched    int                   `json:"replaysWatched"`
	PlayCount         int                   `json:"playCount"`
	LastTryTime       int64                 `json:"lastTryTime"`
	Priority          int                   `json:"priority"`
	OriginalID        int                   `json:"originalId"`
	Player            BeatLeaderPlayer      `json:"player"`
	ScoreImprovement  ScoreImprovement      `json:"scoreImprovement"`
	RankVoting        any                   `json:"rankVoting"`
	Metadata          any                   `json:"metadata"`
	Offsets           BeatLeaderOffsets     `json:"offsets"`
	SotwNominations   int                   `json:"sotwNominations"`
	Status            int                   `json:"status"`
}

// Implement functions for the BeatLeaderResponse struct so it can become a ScoreMessage interface
func (message *BeatLeaderResponse) GetPlayerId() string {
	return message.Player.ID
}
func (message *BeatLeaderResponse) GetPlayerName() string {
	return message.Player.Name
}
func (message *BeatLeaderResponse) GetLeaderboardId() string {
	return message.LeaderboardID
}
func (message *BeatLeaderResponse) GetLeaderboardName() string {
	return message.Leaderboard.Song.Name
}
func (message *BeatLeaderResponse) GetDifficulty() string {
	return message.Leaderboard.Difficulty.DifficultyName
}
func (message *BeatLeaderResponse) GetCountry() string {
	return message.Player.Country
}
func (message *BeatLeaderResponse) GetScore() int {
	return message.ModifiedScore
}
func (message *BeatLeaderResponse) GetPlatform() int {
	return BeatleaderPlatform
}
func (message *BeatLeaderResponse) IsRanked() bool {
	return message.PP > 0
}
func (message *BeatLeaderResponse) GetScoreId() string {
	return strconv.Itoa(message.ID)
}
func (message *BeatLeaderResponse) GetMaxScore() int {
	return message.Leaderboard.Difficulty.MaxScore
}
func (message *BeatLeaderResponse) GetTimestamp() int64 {
	return message.Timepost * 1000
}
func (message *BeatLeaderResponse) GetModifiers() string {
	return message.Modifiers
}
func (message *BeatLeaderResponse) GetBadCuts() int {
	return message.BadCuts
}
func (message *BeatLeaderResponse) GetMissedNotes() int {
	return message.MissedNotes
}

type ContextExtension struct {
	ID               int              `json:"id"`
	PlayerID         string           `json:"playerId"`
	Weight           float64          `json:"weight"`
	Rank             int              `json:"rank"`
	BaseScore        int              `json:"baseScore"`
	ModifiedScore    int              `json:"modifiedScore"`
	Accuracy         float64          `json:"accuracy"`
	PP               float64          `json:"pp"`
	PassPP           float64          `json:"passPP"`
	AccPP            float64          `json:"accPP"`
	TechPP           float64          `json:"techPP"`
	BonusPp          float64          `json:"bonusPp"`
	Modifiers        string           `json:"modifiers"`
	Context          int              `json:"context"`
	ScoreImprovement ScoreImprovement `json:"scoreImprovement"`
}

type ScoreImprovement struct {
	ID                    int     `json:"id"`
	Timeset               string  `json:"timeset"`
	Score                 int     `json:"score"`
	Accuracy              float64 `json:"accuracy"`
	PP                    float64 `json:"pp"`
	BonusPp               float64 `json:"bonusPp"`
	Rank                  int     `json:"rank"`
	AccRight              float64 `json:"accRight"`
	AccLeft               float64 `json:"accLeft"`
	AverageRankedAccuracy float64 `json:"averageRankedAccuracy"`
	TotalPp               float64 `json:"totalPp"`
	TotalRank             int     `json:"totalRank"`
	BadCuts               int     `json:"badCuts"`
	MissedNotes           int     `json:"missedNotes"`
	BombCuts              int     `json:"bombCuts"`
	WallsHit              int     `json:"wallsHit"`
	Pauses                int     `json:"pauses"`
	Modifiers             string  `json:"modifiers"`
}

type BeatLeaderLeaderboard struct {
	ID         string               `json:"id"`
	Song       BeatLeaderSong       `json:"song"`
	Difficulty BeatLeaderDifficulty `json:"difficulty"`
}

type BeatLeaderSong struct {
	ID              string  `json:"id"`
	Hash            string  `json:"hash"`
	Name            string  `json:"name"`
	SubName         string  `json:"subName"`
	Author          string  `json:"author"`
	Mapper          string  `json:"mapper"`
	MapperID        int     `json:"mapperId"`
	CollaboratorIds any     `json:"collaboratorIds"`
	CoverImage      string  `json:"coverImage"`
	BPM             float64 `json:"bpm"`
	Duration        float64 `json:"duration"`
	FullCoverImage  any     `json:"fullCoverImage"`
	Explicity       int     `json:"explicity"`
}

type BeatLeaderDifficulty struct {
	ID              int                       `json:"id"`
	Value           int                       `json:"value"`
	Mode            int                       `json:"mode"`
	DifficultyName  string                    `json:"difficultyName"`
	ModeName        string                    `json:"modeName"`
	Status          int                       `json:"status"`
	ModifierValues  BeatLeaderModifierValues  `json:"modifierValues"`
	ModifiersRating BeatLeaderModifiersRating `json:"modifiersRating"`
	NominatedTime   int64                     `json:"nominatedTime"`
	QualifiedTime   int64                     `json:"qualifiedTime"`
	RankedTime      int64                     `json:"rankedTime"`
	Stars           float64                   `json:"stars"`
	PredictedAcc    float64                   `json:"predictedAcc"`
	PassRating      float64                   `json:"passRating"`
	AccRating       float64                   `json:"accRating"`
	TechRating      float64                   `json:"techRating"`
	Type            int                       `json:"type"`
	Njs             float64                   `json:"njs"`
	Nps             float64                   `json:"nps"`
	Notes           int                       `json:"notes"`
	Bombs           int                       `json:"bombs"`
	Walls           int                       `json:"walls"`
	MaxScore        int                       `json:"maxScore"`
	Duration        float64                   `json:"duration"`
	Requirements    int                       `json:"requirements"`
}

type BeatLeaderModifierValues struct {
	ModifierID int     `json:"modifierId"`
	Da         float64 `json:"da"`
	Fs         float64 `json:"fs"`
	Sf         float64 `json:"sf"`
	Ss         float64 `json:"ss"`
	Gn         float64 `json:"gn"`
	Na         float64 `json:"na"`
	Nb         float64 `json:"nb"`
	Nf         float64 `json:"nf"`
	No         float64 `json:"no"`
	Pm         float64 `json:"pm"`
	Sc         float64 `json:"sc"`
	Sa         float64 `json:"sa"`
	Op         float64 `json:"op"`
	Ez         float64 `json:"ez"`
	Hd         float64 `json:"hd"`
	Smc        float64 `json:"smc"`
	Ohp        float64 `json:"ohp"`
}

type BeatLeaderModifiersRating struct {
	ID              int     `json:"id"`
	SsPredictedAcc  float64 `json:"ssPredictedAcc"`
	SsPassRating    float64 `json:"ssPassRating"`
	SsAccRating     float64 `json:"ssAccRating"`
	SsTechRating    float64 `json:"ssTechRating"`
	SsStars         float64 `json:"ssStars"`
	FsPredictedAcc  float64 `json:"fsPredictedAcc"`
	FsPassRating    float64 `json:"fsPassRating"`
	FsAccRating     float64 `json:"fsAccRating"`
	FsTechRating    float64 `json:"fsTechRating"`
	FsStars         float64 `json:"fsStars"`
	SfPredictedAcc  float64 `json:"sfPredictedAcc"`
	SfPassRating    float64 `json:"sfPassRating"`
	SfAccRating     float64 `json:"sfAccRating"`
	SfTechRating    float64 `json:"sfTechRating"`
	SfStars         float64 `json:"sfStars"`
	BfsPredictedAcc float64 `json:"bfsPredictedAcc"`
	BfsPassRating   float64 `json:"bfsPassRating"`
	BfsAccRating    float64 `json:"bfsAccRating"`
	BfsTechRating   float64 `json:"bfsTechRating"`
	BfsStars        float64 `json:"bfsStars"`
	BsfPredictedAcc float64 `json:"bsfPredictedAcc"`
	BsfPassRating   float64 `json:"bsfPassRating"`
	BsfAccRating    float64 `json:"bsfAccRating"`
	BsfTechRating   float64 `json:"bsfTechRating"`
	BsfStars        float64 `json:"bsfStars"`
}

type BeatLeaderPlayer struct {
	ID                string                    `json:"id"`
	Name              string                    `json:"name"`
	Platform          string                    `json:"platform"`
	Avatar            string                    `json:"avatar"`
	Country           string                    `json:"country"`
	Alias             any                       `json:"alias"`
	Bot               bool                      `json:"bot"`
	Temporary         bool                      `json:"temporary"`
	Pp                float64                   `json:"pp"`
	Rank              int                       `json:"rank"`
	CountryRank       int                       `json:"CountryRank"`
	Level             int                       `json:"level"`
	Experience        int                       `json:"experience"`
	Prestige          int                       `json:"prestige"`
	Role              string                    `json:"role"`
	Socials           any                       `json:"socials"`
	ContextExtensions any                       `json:"contextExtensions"`
	PatreonFeatures   any                       `json:"patreonFeatures"`
	ProfileSettings   BeatLeaderProfileSettings `json:"profileSettings"`
	ClanOrder         string                    `json:"clanOrder"`
	Clans             []any                     `json:"clans"`
}

type BeatLeaderProfileSettings struct {
	ID                    int     `json:"id"`
	Bio                   any     `json:"bio"`
	Message               any     `json:"message"`
	EffectName            string  `json:"effectName"`
	ProfileAppearance     string  `json:"profileAppearance"`
	Hue                   int     `json:"hue"`
	Saturation            float64 `json:"saturation"`
	LeftSaberColor        any     `json:"leftSaberColor"`
	RightSaberColor       any     `json:"rightSaberColor"`
	ProfileCover          string  `json:"profileCover"`
	StarredFriends        string  `json:"starredFriends"`
	HorizontalRichBio     bool    `json:"horizontalRichBio"`
	RankedMapperSort      any     `json:"rankedMapperSort"`
	ShowBots              bool    `json:"showBots"`
	ShowAllRatings        bool    `json:"showAllRatings"`
	ShowExplicitCovers    bool    `json:"showExplicitCovers"`
	ShowStatsPublic       bool    `json:"showStatsPublic"`
	ShowStatsPublicPinned bool    `json:"showStatsPublicPinned"`
}

type BeatLeaderOffsets struct {
	ID           int `json:"id"`
	Frames       int `json:"frames"`
	Notes        int `json:"notes"`
	Walls        int `json:"walls"`
	Heights      int `json:"heights"`
	Pauses       int `json:"pauses"`
	SaberOffsets int `json:"saberOffsets"`
	CustomData   int `json:"customData"`
}
