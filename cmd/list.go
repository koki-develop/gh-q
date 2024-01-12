package cmd

import (
	"fmt"

	"github.com/koki-develop/gh-q/internal/cli"
	"github.com/koki-develop/go-fzf"
	"github.com/spf13/cobra"
)

var (
	flagListFullPath bool
	flagListFilter   bool
	flagListMultiple bool
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List managed repositories",
	Long:    "List managed repositories.",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cli.NewClient()
		if err != nil {
			return err
		}

		dirs, err := c.ListDirectories()
		if err != nil {
			return err
		}

		items := make([]string, len(dirs))
		if flagListFullPath {
			for i, d := range dirs {
				items[i] = d.FullPath
			}
		} else {
			for i, d := range dirs {
				items[i] = fmt.Sprintf("%s/%s", d.Owner, d.Repo)
			}
		}

		if flagListFilter {
			opts := []fzf.Option{}
			if flagListMultiple {
				opts = append(opts, fzf.WithNoLimit(true))
			}

			f, err := fzf.New(opts...)
			if err != nil {
				return err
			}

			idxs, err := f.Find(items, func(i int) string {
				return items[i]
			})
			if err != nil {
				return err
			}
			for _, idx := range idxs {
				fmt.Println(items[idx])
			}

			return nil
		}

		for _, item := range items {
			fmt.Println(item)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&flagListFullPath, "full-path", "p", false, "print full path")
	listCmd.Flags().BoolVarP(&flagListFilter, "filter", "f", false, "filter by fuzzy search")
	listCmd.Flags().BoolVarP(&flagListMultiple, "multiple", "m", false, "allow multiple selection (only available with --filter)")
}
