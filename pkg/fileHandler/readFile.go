/*
This package is responsible for reading the file provided by user
  - if the file type is supported it runs the schema validators and validates the file
    if the file is not of valid schema , it prints the error and exits.
  - if both checks pass it returns the questions
*/
package fileHandler

import (
	"fmt"
	"os"
	"quiz/pkg/colors"
	"quiz/pkg/customErrors"
	"quiz/pkg/types"
	"quiz/pkg/utils"
	"quiz/pkg/validators"
	"strconv"
)

// GetQuestions return a slice of Problem type and return InvalidFileType/InvalidSchema error if file type or schema is invalid
func GetQuestions(args []string) ([]types.Problem, error) {
	fname := args[0] // file name
	oFile, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer oFile.Close()
	problems, err := validators.IsValid(oFile)
	if err != nil {
		if utils.PrintErrorsAndWarnings() {
			if schemaErrors, ok := err.(*customErrors.ErrInvalidProblems); ok {
				for _, reason := range schemaErrors.InvalidQuestions {
					for qno, r := range reason {
						q := strconv.Itoa(qno)
						fmt.Println(
							colors.GraveError.Render("Disregarding question "+q+" because :\n",
								colors.NormalText.Render(r)))
					}
				}
				for _, reason := range schemaErrors.Warnings {
					for qno, r := range reason {
						q := strconv.Itoa(qno)
						fmt.Println(
							colors.MuchWarning.Render("Warning for Question "+q+" : \n",
								colors.NormalText.Render(r)))
					}
				}
			} else {
				return nil, err
			}
		}
	}
	if len(problems) == 0 {
		return nil, customErrors.ErrInvalidSchema
	}
	return problems, nil
}
