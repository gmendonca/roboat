package github

import (
	"github.com/google/go-github/github"
)

func CommentOnGithubIssuet(client *github.Client, payload GithubIssueCommentPayload, body string) error {
	_, _, err := client.Issues.CreateComment(
		payload.Repository.Owner.Login,
		payload.Repository.Name,
		payload.Issue.Number,
		&github.IssueComment{Body: &body},
	)
	return err
}
