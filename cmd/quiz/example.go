package quiz

import (
	"quiz/pkg/examples"
	"quiz/pkg/utils"

	"github.com/spf13/cobra"
)

var generateExamples = &cobra.Command{
	Use:     "example",
	Aliases: []string{"eg"},
	Short:   "Generates an example.json file with couple questions for reference",
	Long:    "Generates an examples.json file that has 3 questions with valid schema",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := examples.GenExamples()
		if err != nil {
			utils.ExitWithMessage("Something went wrong", 1)
		}

	},
}

func init() {
	rootCmd.AddCommand(generateExamples)
}
