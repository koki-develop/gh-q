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

		if flagListFilter {
			opts := []fzf.Option{}
			if flagListMultiple {
				opts = append(opts, fzf.WithNoLimit(true))
			}

			f, err := fzf.New(opts...)
			if err != nil {
				return err
			}

			idxs, err := f.Find(dirs, func(i int) string {
				return dirs[i].Path(flagListFullPath)
			})
			if err != nil {
				return err
			}
			for _, idx := range idxs {
				fmt.Println(dirs[idx].Path(flagListFullPath))
			}

			return nil
		}

		for _, dir := range dirs {
			fmt.Println(dir.Path(flagListFullPath))
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
