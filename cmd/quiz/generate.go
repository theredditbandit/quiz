package quiz

import (
	"fmt"
	"quiz/pkg/fileHandler"
	"quiz/pkg/utils"
	"strconv"

	"github.com/spf13/cobra"
)

var name string

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen", "g"},
	Short:   "Generates boilerplate JSON for MCQ type questions",
	Long:    "Takes the number of questions to generate and generates a problems.json file",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			utils.ExitWithMessage(fmt.Sprintf("Invalid argument %s: argument must be an integer", args[0]), 1)
		}

		err = fileHandler.GenBoilerplate(n, name)
		if err != nil {
			utils.ExitWithMessage(fmt.Sprintf("Failed to generate boilerplate: %s", err.Error()), 1)
		}

	},
}

func init() {
	rootCmd.Flags().StringVar(&name, "name", "problems", "Specify the name of the output file")
	rootCmd.AddCommand(generateCmd)
}
