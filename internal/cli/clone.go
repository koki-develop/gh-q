package cli

import (
	"path/filepath"

	"github.com/koki-develop/gh-q/internal/git"
)

func (c *Client) Clone(owner, repo string) error {
	dest := filepath.Join(c.root, "github.com", owner, repo)

	opts := []git.CloneOption{}
	auth, err := c.auth()
	if err != nil {
		return err
	}
	if auth != nil {
		opts = append(opts, git.WithCloneAuth(auth.Username, auth.Token))
	}

	if err := c.gitClient.Clone(owner, repo, dest, opts...); err != nil {
		return err
	}
	return nil
}
