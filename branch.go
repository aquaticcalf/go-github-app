package gogithubapp

import (
	"context"

	githubapi "github.com/google/go-github/v50/github"
)

// CreateBranch creates a new branch from baseBranch
func (c *Client) CreateBranch(ctx context.Context, owner, repo, branchName, baseBranch string) (*githubapi.Reference, error) {
	baseRef, _, err := c.client.Git.GetRef(ctx, owner, repo, "refs/heads/"+baseBranch)
	if err != nil {
		return nil, err
	}
	ref := &githubapi.Reference{Ref: githubapi.String("refs/heads/" + branchName), Object: &githubapi.GitObject{SHA: baseRef.Object.SHA}}
	apiref, _, err := c.client.Git.CreateRef(ctx, owner, repo, ref)
	if err != nil {
		return nil, err
	}
	return apiref, nil
}

// ListBranches lists all branches in a repository
func (c *Client) ListBranches(ctx context.Context, owner, repo string) ([]*githubapi.Branch, error) {
	var all []*githubapi.Branch
	opt := &githubapi.BranchListOptions{ListOptions: githubapi.ListOptions{PerPage: 50}}
	for {
		branches, resp, err := c.client.Repositories.ListBranches(ctx, owner, repo, opt)
		if err != nil {
			return nil, err
		}
		all = append(all, branches...)
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}
	return all, nil
}
