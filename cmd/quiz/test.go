package quiz

import (
	"fmt"
	"quiz/pkg/customErrors"
	"quiz/pkg/fileHandler"
	"quiz/pkg/testUser"
	"quiz/pkg/utils"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"tst", "t"},
	Short:   "Takes a file containing the test questions as an argument.",
	Long:    "Takes a JSON file of format \nsee 'quiz help generate' or 'quiz help example' for more information",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		questions, err := fileHandler.GetQuestions(args)
		if err != nil {
			switch err {
			case customErrors.ErrInvalidSchema:
				utils.ExitWithMessage("Invalid Schema: Schema of provided file is not valid.", 1)
			case customErrors.ErrInvalidFileType:
				utils.ExitWithMessage("Invalid File type: only JSON file format is supported.", 1)
				cmd.Help()
			default:
				utils.ExitWithMessage(fmt.Sprintf("Unknown Error: %s", err), 1)
			}
		}
		testUser.QuestionUser(questions)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
