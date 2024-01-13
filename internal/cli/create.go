package cli

import (
	"errors"
	"os"
)

func (c *Client) Create(owner, repo string) error {
	p := c.path(owner, repo)
	if _, err := os.Stat(p); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err := os.MkdirAll(p, 0755); err != nil {
			return err
		}
	}

	ok, err := c.gitClient.IsExists(p)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("already exists")
	}

	if err := c.gitClient.Init(p); err != nil {
		return err
	}

	return nil
}
