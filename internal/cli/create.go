package cli

import "os"

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

	if err := c.gitClient.Init(p); err != nil {
		return err
	}

	return nil
}
