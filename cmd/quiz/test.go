package quiz

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"quiz/pkg/customErrors"
	"quiz/pkg/customTypes"
	"quiz/pkg/fileHandler"
	"quiz/pkg/ui"
	"quiz/pkg/utils"
)

var (
	time    int
	shuffle bool
	min     bool
	hour    bool
	sec     bool
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Takes a file containing the test questions as an argument",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		questions, err := filehandler.GetQuestions(args)
		if err != nil {

			switch err {
			case customerrors.InvalidSchemaError:
				utils.ExitWithMessage("Invalid Schema: Schema of provided file is not valid.", 1)
			case customerrors.InvalidFileTypeError:
				utils.ExitWithMessage("Invalid File type: only CSV and JSON file formats are supported.", 1)
				cmd.Help()
			}
		}
		if shuffle {
			rand.Shuffle(len(questions), func(i, j int) {
				questions[i], questions[j] = questions[j], questions[i]
			})
		}

		timeConf := handleTimeConf(time, sec, min, hour)
		marks, err := ui.QuestionUser(questions, timeConf, ui.ConsoleReader, ui.QuizTimer)
		printMarksHandleErrors(marks, err, questions)
	},
}

func init() {
	testCmd.Flags().BoolVar(&shuffle, "shuffle", false, "Whether or not to shuffle the questions")
	testCmd.Flags().IntVar(&time, "time", 0, "Time limit for the quiz , defaults to untimed quiz") // TODO: make this  configuratble via config file in the future.

	rootCmd.AddCommand(testCmd)
}

func printMarksHandleErrors(marks int, errors error, problems []customTypes.Problem) {
	fmt.Printf("You got %d/%d correct!\n", marks, len(problems))
	if errors != nil {
		userErrs, _ := errors.(ui.QuizEvaluation)
		userErrs.PrintErrors()
		if userErrs.Unattempted {
			userErrs.PrintUnattempted()
		}
	}
}

func handleTimeConf(time int, sec bool, min bool, hour bool) customTypes.TimeConf {
	tconf := customTypes.TimeConf{Time: time, Unit: "sec"}

	if min {
		tconf.Unit = "min"
	} else if hour {
		tconf.Unit = "hour"
	}
	return tconf
}
