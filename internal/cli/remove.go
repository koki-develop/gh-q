package cli

import (
	"fmt"
	"os"

	"github.com/koki-develop/gh-q/internal/util"
)

type removeOptions struct {
	Force bool
}

type RemoveOption func(*removeOptions)

func WithRemoveForce(b bool) RemoveOption {
	return func(o *removeOptions) {
		o.Force = b
	}
}

func (c *Client) Remove(owner, repo string, opts ...RemoveOption) error {
	o := &removeOptions{}
	for _, opt := range opts {
		opt(o)
	}

	p := c.path(owner, repo)

	if !o.Force {
		if ok := util.Confirm(fmt.Sprintf("Remove `%s`?", p)); !ok {
			fmt.Println("Canceled.")
			return nil
		}
	}

	if err := os.RemoveAll(p); err != nil {
		return err
	}
	fmt.Println("Removed.")

	return nil
}
