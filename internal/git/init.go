package git

import "github.com/go-git/go-git/v5"

func (c *Client) Init(p string) error {
	if _, err := git.PlainInit(p, false); err != nil {
		return err
	}
	return nil
}
