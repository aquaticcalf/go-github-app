package gogithubapp

import (
	"context"

	githubapi "github.com/google/go-github/v50/github"
)

// CreateIssue opens a new issue
func (c *Client) CreateIssue(ctx context.Context, owner, repo, title, body string) (*githubapi.Issue, error) {
	req := &githubapi.IssueRequest{Title: githubapi.String(title), Body: githubapi.String(body)}
	issue, _, err := c.client.Issues.Create(ctx, owner, repo, req)
	if err != nil {
		return nil, err
	}
	return issue, nil
}

// ListIssues lists issues for a repo
func (c *Client) ListIssues(ctx context.Context, owner, repo string, opts *githubapi.IssueListByRepoOptions) ([]*githubapi.Issue, error) {
	issues, _, err := c.client.Issues.ListByRepo(ctx, owner, repo, opts)
	return issues, err
}
