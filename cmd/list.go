package cmd

import (
	"fmt"

	"github.com/koki-develop/gh-q/internal/cli"
	"github.com/spf13/cobra"
)

var (
	flagListFullPath bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List managed repositories",
	Long:  "List managed repositories.",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cli.NewClient()
		if err != nil {
			return err
		}

		dirs, err := c.ListDirectories()
		if err != nil {
			return err
		}

		if flagListFullPath {
			for _, d := range dirs {
				fmt.Println(d.FullPath)
			}
			return nil
		}

		for _, d := range dirs {
			fmt.Printf("%s/%s\n", d.Owner, d.Repo)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&flagListFullPath, "full-path", "p", false, "print full path")
}
