package cli

import (
	"os"
	"path/filepath"
)

type Directory struct {
	Owner    string
	Repo     string
	FullPath string
}

func (d *Directory) Path(full bool) string {
	if full {
		return d.FullPath
	} else {
		return filepath.Join("github.com", d.Owner, d.Repo)
	}
}

func (c *Client) ListDirectories() ([]*Directory, error) {
	matches, err := filepath.Glob(filepath.Join(c.root, "github.com/*/*"))
	if err != nil {
		return nil, err
	}

	dirs := []*Directory{}
	for _, match := range matches {
		info, err := os.Stat(match)
		if err != nil {
			return nil, err
		}
		if !info.IsDir() {
			continue
		}

		ok, err := c.gitClient.IsExists(match)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		dirs = append(dirs, &Directory{
			Owner:    filepath.Base(filepath.Dir(match)),
			Repo:     filepath.Base(match),
			FullPath: match,
		})
	}

	return dirs, nil
}
