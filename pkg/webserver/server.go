package webserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Webserver struct{
	Port string
}

//StartServer start a Gin Tonic Server to
func (webserver *Webserver) StartServer() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := router.Group("/api"){
		api.POST("/github", handleGitHubRequests)
	}
	
	router.Run(webserver.Port)
}

func (webserver *Webserver) handleGitHubRequests(c *gin.Context) {
	if c.Request.Header.Get("X-GitHub-Event") == "issue_comment" {
		var json GitHubIssueCommentPayload
		if err:= c.BindJSON(&json}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.Action != "created" {
			c.JSON(http.StatusOK, gin.H{"status:" "Skipping event"})
			return
		}

		if err := handleGitHubIssueCommentPayload(json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Event handled"})
	} else {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

func (webserver *Webserver) handleGitHubIssueCommentPayload(payload GitHubIssueCommentPayload) error {
	
}
