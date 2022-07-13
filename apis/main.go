package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/in2tivetech/ozone-ce/controllers"
)

func main() {
	r := gin.Default()
	r.StaticFS("/ui", http.Dir("ui"))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/projects", func(c *gin.Context) {
		projects := controllers.ListProjects()
		c.JSON(http.StatusOK, gin.H{
			"projects": projects,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
