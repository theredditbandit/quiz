/*
This package is responsible for reading the file provided by user
  - determinig the file type and checking if the file type is supported (csv,json).
  - if the file type is supported it runs the schema validators and validates the file
    if the file is not of valid schema , it panics and exit.
  - if both checks pass it returns the questions
*/
package filehandler

import (
	"encoding/csv"
	"fmt"
	"os"
	customerrors "quiz/pkg/customErrors"
	"quiz/pkg/parsers"
	"quiz/pkg/types"
	"quiz/pkg/utils"
	"quiz/pkg/validators"
)

// GetQuestions from the file
func GetQuestions(args []string) ([]types.Problem, error) {
	file := args[0]
	fileType := utils.GetFileType(file)

	if fileType == "csv" || fileType == "json" {
		if validators.IsValid(file) {
			oFile, err := os.Open(file)
			if err != nil {
				utils.ExitWithMessage(fmt.Sprintf("Failed to open the CSV file: %s\n", file), 1)
			}
			defer oFile.Close()
			reader := csv.NewReader(oFile)
			lines, err := reader.ReadAll()

			if err != nil {
				utils.ExitWithMessage("Something went wrong while reading lines from the csv", 1)
			}
			return parsers.LinesToProblems(lines), nil
		} else {
			return nil, customerrors.InvalidSchemaError
		}
	} else {
		return nil, customerrors.InvalidFileTypeError
	}
}
