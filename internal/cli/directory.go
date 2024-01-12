package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

type Directory struct {
	Name     string
	FullPath string
}

func (d *Directory) Path(full bool) string {
	if full {
		return d.FullPath
	} else {
		return fmt.Sprintf(d.Name)
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

		name, err := filepath.Rel(c.root, p)
		if err != nil {
			return err
		}

		dirs = append(dirs, &Directory{
			Name:     name,
			FullPath: p,
		})
		return filepath.SkipDir
	})
	if err != nil {
		return nil, err
	}

	return dirs, nil
}
