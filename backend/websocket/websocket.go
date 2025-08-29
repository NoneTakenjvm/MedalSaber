package websocket

import (
	"fmt"
	"time"

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
	for {
		// Connect to the socket
		c, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			fmt.Printf("Error connecting to socket %s: %v\n", url, err)
			// Wait before retrying
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf("Connected to %s\n", url)

		// Close the socket when the function returns
		defer c.Close()

		// Read messages from the server
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Printf("Error reading message from %s: %v\n", url, err)
				break
			}
			callback(message)
		}
		fmt.Printf("Connection lost to %s, reconnecting...\n", url)

		// Wait before reconnecting
		time.Sleep(5 * time.Second)
	}
}
