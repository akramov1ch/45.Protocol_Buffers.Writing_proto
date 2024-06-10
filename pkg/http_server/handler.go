package http_server

import (
	"45HW/pkg/http_server/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/request", handleRequest)
	return router
}

func handleRequest(c *gin.Context) {
	word := c.Query("word")
	result, err := client.SendRequest(word)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}
