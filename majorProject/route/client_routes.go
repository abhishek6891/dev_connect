package routes

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/user/data/client"
)

func RegisterClientRoutes(route *gin.Engine) {
	route.GET("/clients", func(c *gin.Context) {
		filePath := "Doc/project/client.json"
		clientsList, err := client.LoadClientsFromFile(filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": clientsList})
	})
}
