package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// browser - view
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		//deploy app with the given details
		c.JSON(http.StatusOK, gin.H{
			"message": GetTheMessage(),
		})
	})

	r.GET("/world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "world",
		})
	})

	r.GET("/someGet", getting)
	r.POST("/somePost", posting)
	// r.PUT("/somePut", putting)
	// r.DELETE("/someDelete", deleting)
	// r.PATCH("/somePatch", patching)
	// r.HEAD("/someHead", head)
	// r.OPTIONS("/someOptions", options)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "getting",
	})
}

func posting(c *gin.Context) {
	//database manipulation

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully posted data",
	})
}

// controller
func GetTheMessage() string {
	return "Hellow World!!"
}

// cli - view
func handleCliPing() {
	fmt.Println(GetTheMessage())
}
