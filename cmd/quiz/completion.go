package quiz

import "github.com/spf13/cobra"

func completionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}
}

func init() {
	completion := completionCmd()
	completion.Hidden = true
	rootCmd.AddCommand(completion)
}
