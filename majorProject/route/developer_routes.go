package routes

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/user/data/developer"
)

func RegisterDeveloperRoutes(route *gin.Engine) {
	route.GET("/developers", func(c *gin.Context) {
		filePath := "Doc/project/developer.json"
		developersList, err := developer.LoadDevelopersFromFile(filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": developersList})
	})
}
