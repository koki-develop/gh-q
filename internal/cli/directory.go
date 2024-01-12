package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
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
		return fmt.Sprintf("%s/%s", d.Owner, d.Repo)
	}
}

func (c *Client) ListDirectories() ([]*Directory, error) {
	ps, err := filepath.Glob(filepath.Join(c.root, "github.com/*/*"))
	if err != nil {
		return nil, err
	}

	dirs := []*Directory{}
	for _, p := range ps {
		info, err := os.Stat(p)
		if err != nil {
			return nil, err
		}

		// check if directory
		if !info.IsDir() {
			continue
		}

		// check if git repository
		if _, err := git.PlainOpen(p); err != nil {
			if err == git.ErrRepositoryNotExists {
				continue
			}
			return nil, err
		}

		dirs = append(dirs, &Directory{
			Owner:    filepath.Base(filepath.Dir(p)),
			Repo:     filepath.Base(p),
			FullPath: p,
		})
	}

	return dirs, nil

}
