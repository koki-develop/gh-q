package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Init(p string) error {
	if _, err := git.PlainInit(p, false); err != nil {
		return err
	}
	return nil
}

type cloneOptions struct {
	Auth *Auth
}

type CloneOption func(*cloneOptions)

func WithCloneAuth(username, token string) CloneOption {
	return func(opts *cloneOptions) {
		opts.Auth = &Auth{
			Username: username,
			Token:    token,
		}
	}
}

func (c *Client) Clone(owner, repo, dest string, opts ...CloneOption) error {
	o := &cloneOptions{}
	for _, opt := range opts {
		opt(o)
	}

	copts := &git.CloneOptions{
		URL: fmt.Sprintf("https://github.com/%s/%s", owner, repo),
	}
	if o.Auth != nil {
		copts.Auth = &http.BasicAuth{
			Username: o.Auth.Username,
			Password: o.Auth.Token,
		}
	}

	fmt.Printf("Cloning into %s\n", dest)
	if _, err := git.PlainClone(dest, false, copts); err != nil {
		return err
	}
	fmt.Println("Done.")

	return nil
}
