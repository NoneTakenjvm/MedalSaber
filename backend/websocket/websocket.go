package websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
	"nonetaken.dev/medalsaber/score"
)

func Initialise() {
	// Connect to the WebSocket server as a client
	c, _, err := websocket.DefaultDialer.Dial("wss://scoresaber.com/ws", nil)
	if err != nil {
		fmt.Printf("dial error: %v", err)
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
		score.HandleScore(message)
	}
}
