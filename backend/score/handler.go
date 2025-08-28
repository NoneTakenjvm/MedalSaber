package score

import (
	"encoding/json"
	"fmt"
	"log"
)

func HandleScore(message []byte) {
	var incomingScoreMessage IncomingMessageWithScore
	err := json.Unmarshal([]byte(message), &incomingScoreMessage)
	if err != nil {
		log.Printf("Error while parsing incoming message: %s\n", err)
		return
	}
	incomingScore := incomingScoreMessage.Score
	fmt.Println(incomingScore.Leaderboard.SongName)
}
