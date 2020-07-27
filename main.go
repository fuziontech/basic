package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/posthog/posthog-go"
)

func main() {
	r := gin.Default()
	client := posthog.New("Y7NloBxRA3R86YIMEzT1rpJ9yFsCc_qiftFeWNFM9lA")
	defer client.Close()

	r.GET("/", func(c *gin.Context) {
		// Capture an event
		client.Enqueue(posthog.Capture{
			DistinctId: "test-user",
			Event:      "pageview",
			Properties: posthog.NewProperties().
				Set("app", "basic"),
		})
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/health" func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
	r.Run()
}
