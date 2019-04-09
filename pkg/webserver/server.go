package webserver

import "github.com/gin-gonic/gin"

//StartServer start a Gin Tonic Server to
func StartServer() {
	router = gin.Default()

	api := router.Group("/api")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	webserver.Server.Run()
}
