package quiz

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "quiz",
	Short:   "quiz - a simple quiz system for the terminal",
	Long:    "A quiz application that reads the questions from a CSV or JSON file,with support for Question andwer type problems or MCQ problems",
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Somethig went wrong '%s'\n", err)
		os.Exit(1)
	}
}
