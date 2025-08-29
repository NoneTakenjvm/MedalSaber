package websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
	"nonetaken.dev/medalsaber/score"
)

func Initialise() {
	// Initialise ScoreSaber
	go func() {
		initSocket("wss://scoresaber.com/ws", func(message []byte) {
			score.HandleScore(score.ScoresaberPlatform, message)
		})
		fmt.Println("Connected to ScoreSaber")
	}()
	go func() {
		// Initialise BeatLeader
		initSocket("wss://sockets.api.beatleader.com/scores", func(message []byte) {
			score.HandleScore(score.BeatleaderPlatform, message)
		})
		fmt.Println("Connected to BeatLeader")
	}()
}

func initSocket(url string, callback func(message []byte)) {
	// Connect to the Scoresaber socket
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Printf("Error connecting to socket: %v", err)
		return
	}
	// Close the socket when the function returns
	defer c.Close()

	// Read messages from the server
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			continue
		}
		callback(message)
	}
}
