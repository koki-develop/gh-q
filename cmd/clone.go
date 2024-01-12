package cmd

import "github.com/spf13/cobra"

var cloneCmd = &cobra.Command{
	Use:  "clone",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
