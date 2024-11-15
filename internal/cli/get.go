package cli

import (
	"github.com/koki-develop/gh-q/internal/git"
)

func (c *Client) Get(owner, repo string) error {
	dest := c.path(owner, repo)

	opts := []git.CloneOption{}

	sshKeyPath, err := GetSSHKeyPath()
	if err != nil {
		return err
	}
	if sshKeyPath != "" {
		opts = append(opts, git.WithCloneSSHKey(sshKeyPath))
	} else {
		auth, err := c.auth()
		if err != nil {
			return err
		}
		if auth != nil {
			opts = append(opts, git.WithCloneAuth(auth.Username, auth.Token))
		}
	}

	if err := c.gitClient.Clone(owner, repo, dest, opts...); err != nil {
		return err
	}
	return nil
}
