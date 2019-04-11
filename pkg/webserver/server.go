package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmendonca/roboat/pkg/roboat"
)

// Webserver struct
type Webserver struct {
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

	api := router.Group("/api")
	{
		api.POST("/github", webserver.HandleGitHubIssueCommentRequests)
	}

	router.Run(webserver.Port)
}

// HandleGitHubIssueCommentRequests handles requests from GitHub Issue Comment Requests
func (webserver *Webserver) HandleGitHubIssueCommentRequests(c *gin.Context) {
	if c.Request.Header.Get("X-GitHub-Event") == "issue_comment" {
		var json roboat.GitHubIssueCommentPayload
		if err := c.BindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.Action != "created" {
			c.JSON(http.StatusOK, gin.H{"status": "Skipping event"})
			return
		}

		if err := webserver.HandleGitHubIssueCommentPayload(json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Event handled"})
	} else {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

func (webserver *Webserver) HandleGitHubIssueCommentPayload(payload roboat.GitHubIssueCommentPayload) error {

	return nil
}
