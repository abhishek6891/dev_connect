package routes

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/projects"
)

func RegisterProjectRoutes(route *gin.Engine) {
	route.GET("/projects", func(c *gin.Context) {
		filePath := "Doc/project/project.json"
		projectsList, err := projects.LoadProjectsFromFile(filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": projectsList})
	})
}
