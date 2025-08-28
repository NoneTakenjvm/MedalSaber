package main

import (
	"context"
	"fmt"

	"nonetaken.dev/medalsaber/database"
)

func main() {
	// Initialise the database handler
	database.Initialise()
	fmt.Println("Database initialised")
	// Define a defer function to handle the client disconnecting
	defer func() {
		if err := database.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
