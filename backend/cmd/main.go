package main

import (
	"context"
	"fmt"

	"nonetaken.dev/medalsaber/database"
	"nonetaken.dev/medalsaber/websocket"
)

func main() {
	// Initialise the database handler
	database.Initialise()
	fmt.Println("Database initialised")

	// Initialise the websocket handler
	websocket.Initialise()
	fmt.Println("Websocket handler initialised")

	// Define a defer function to handle the client disconnecting
	defer func() {
		if err := database.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	select {}
}
