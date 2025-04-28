package gogithubapp

import (
	"context"

	githubapi "github.com/google/go-github/v50/github"
)

// CommitMultipleFiles commits many files atomically to a branch
func (c *Client) CommitMultipleFiles(ctx context.Context, owner, repo, branch, message string, files map[string][]byte) (*githubapi.Commit, error) {
	// get base ref
	ref, _, err := c.client.Git.GetRef(ctx, owner, repo, "refs/heads/"+branch)
	if err != nil {
		return nil, err
	}
	baseSHA := *ref.Object.SHA

	// build tree entries
	var entries []*githubapi.TreeEntry
	for path, content := range files {
		blob := &githubapi.Blob{Content: githubapi.String(string(content)), Encoding: githubapi.String("utf-8")}
		bRes, _, err := c.client.Git.CreateBlob(ctx, owner, repo, blob)
		if err != nil {
			return nil, err
		}
		entries = append(entries, &githubapi.TreeEntry{Path: githubapi.String(path), Mode: githubapi.String("100644"), Type: githubapi.String("blob"), SHA: bRes.SHA})
	}

	// create new tree
	tree, _, err := c.client.Git.CreateTree(ctx, owner, repo, baseSHA, entries)
	if err != nil {
		return nil, err
	}

	// create commit
	commit := &githubapi.Commit{Message: githubapi.String(message), Tree: &githubapi.Tree{SHA: tree.SHA}, Parents: []*githubapi.Commit{{SHA: ref.Object.SHA}}}
	newCommit, _, err := c.client.Git.CreateCommit(ctx, owner, repo, commit)
	if err != nil {
		return nil, err
	}

	// update ref
	ref.Object.SHA = newCommit.SHA
	_, _, err = c.client.Git.UpdateRef(ctx, owner, repo, ref, false)
	return newCommit, err
}

// ListCommits lists commits on a branch.
func (c *Client) ListCommits(ctx context.Context, owner, repo string, opts *githubapi.CommitsListOptions) ([]*githubapi.RepositoryCommit, error) {
	commits, _, err := c.client.Repositories.ListCommits(ctx, owner, repo, opts)
	if err != nil {
		return nil, err
	}
	return commits, nil
}

// GetFileContent retrieves file content from a repo.
func (c *Client) GetFileContent(ctx context.Context, owner, repo, path, ref string) (string, *githubapi.RepositoryContent, error) {
	content, _, _, err := c.client.Repositories.GetContents(ctx, owner, repo, path, &githubapi.RepositoryContentGetOptions{Ref: ref})
	if err != nil {
		return "", nil, err
	}
	decoded, err := content.GetContent()
	return decoded, content, err
}
