package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"nonetaken.dev/medalsaber/database"
)

func Initialise() {
	router := gin.Default()

	// Load routes
	router.GET("/player/:platform/:region/:playerId", getPlayer)
	router.GET("/changes/:platform/:region/:playerId", getChanges)

	// Begin the API
	router.Run("localhost:6969")
}

func getPlayer(c *gin.Context) {
	platform, err := strconv.Atoi(c.Param("platform"))
	// Was a correct platform provided?
	if err != nil || (platform != 1 && platform != 2) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid platform, use 1 for ScoreSaber or 2 for Beatleader"})
		return
	}
	// Get the player by the platform
	player, err := database.GetPlayer(platform, c.Param("region"), c.Param("playerId"), "", false)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Player not found"})
		return
	}
	// Return the player
	c.IndentedJSON(http.StatusOK, player)
}

func getChanges(c *gin.Context) {
	// Get the requested platform
	platform, err := strconv.Atoi(c.Param("platform"))
	if err != nil || (platform != 1 && platform != 2) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid platform, use 1 for ScoreSaber or 2 for Beatleader"})
		return
	}
	region := c.Param("region")
	playerId := c.Param("playerId")
	// Parse optional page param
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid page"})
	}
	// Parse optional before param
 	before, err := strconv.ParseInt(c.DefaultQuery("before", "0"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid before"})
	}
	// Parse optional after param
	after, err := strconv.ParseInt(c.DefaultQuery("after", "0"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid after"})
	}
	// Fetch the changes
	changes, err := database.GetChanges(platform, region, playerId, page, before, after)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Changes not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, changes)
}