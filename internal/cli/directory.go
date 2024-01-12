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
		return fmt.Sprintf("github.com/%s/%s", d.Owner, d.Repo)
	}
}

func (c *Client) ListDirectories() ([]*Directory, error) {
	dirs := []*Directory{}
	err := filepath.WalkDir(filepath.Join(c.root, "github.com"), func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}

		if _, err := git.PlainOpen(p); err != nil {
			if err == git.ErrRepositoryNotExists {
				return nil
			}
			return err
		}

		dirs = append(dirs, &Directory{
			Owner:    filepath.Base(filepath.Dir(p)),
			Repo:     filepath.Base(p),
			FullPath: p,
		})
		return filepath.SkipDir
	})
	if err != nil {
		return nil, err
	}

	return dirs, nil
}
