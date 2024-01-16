package quiz

import (
	"fmt"
	"github.com/spf13/cobra"
	"quiz/pkg/utils"
)

var rootCmd = &cobra.Command{
	Use:     "quiz",
	Short:   "quiz - a simple quiz system for the terminal",
	Long:    "A quiz application that reads the questions from a JSON file,with support for Question andwer type problems or MCQ problems",
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			utils.ExitWithMessage(fmt.Sprintf("Something went wrong '%s'\n", err), 1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.ExitWithMessage(fmt.Sprintf("Something went wrong '%s'\n", err), 1)
	}
}
