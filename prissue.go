package gogithubapp

import (
	"context"

	githubapi "github.com/google/go-github/v50/github"
)

// CreateComment adds a comment to an issue or pr
func (c *Client) CreateComment(ctx context.Context, owner, repo string, number int, comment string) (*githubapi.IssueComment, error) {
	req := &githubapi.IssueComment{Body: githubapi.String(comment)}
	createdComment, _, err := c.client.Issues.CreateComment(ctx, owner, repo, number, req)
	if err != nil {
		return nil, err
	}
	return createdComment, nil
}

// AddLabels adds labels to an issue or pr
func (c *Client) AddLabels(ctx context.Context, owner, repo string, number int, labels []string) ([]*githubapi.Label, error) {
	createdLabels, _, err := c.client.Issues.AddLabelsToIssue(ctx, owner, repo, number, labels)
	if err != nil {
		return nil, err
	}
	return createdLabels, nil
}
