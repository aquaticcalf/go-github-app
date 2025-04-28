package gogithubapp

import (
	"context"

	githubapi "github.com/google/go-github/v50/github"
)

// MergePullRequest merges a pull request
func (c *Client) MergePullRequest(ctx context.Context, owner, repo string, number int, commitMessage string) (*githubapi.PullRequestMergeResult, error) {
	mergeResult, _, err := c.client.PullRequests.Merge(ctx, owner, repo, number, commitMessage, nil)
	if err != nil {
		return nil, err
	}
	return mergeResult, nil
}

// CreatePullRequest opens a pull request
func (c *Client) CreatePullRequest(ctx context.Context, owner, repo, title, head, base, body string) (*githubapi.PullRequest, error) {
	req := &githubapi.NewPullRequest{Title: githubapi.String(title), Head: githubapi.String(head), Base: githubapi.String(base), Body: githubapi.String(body)}
	pr, _, err := c.client.PullRequests.Create(ctx, owner, repo, req)
	if err != nil {
		return nil, err
	}
	return pr, nil
}

// ListPullRequests lists pull requests for a repo
func (c *Client) ListPullRequests(ctx context.Context, owner, repo string, opts *githubapi.PullRequestListOptions) ([]*githubapi.PullRequest, error) {
	prs, _, err := c.client.PullRequests.List(ctx, owner, repo, opts)
	return prs, err
}
