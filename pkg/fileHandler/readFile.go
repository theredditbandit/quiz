/*
This package is responsible for reading the file provided by user
  - determinig the file type and checking if the file type is supported (csv,json).
  - if the file type is supported it runs the schema validators and validates the file
    if the file is not of valid schema , it prints the error and exits.
  - if both checks pass it returns the questions
*/
package fileHandler

import (
	"fmt"
	"os"
	"quiz/pkg/customErrors"
	"quiz/pkg/types"
	"quiz/pkg/validators"
)

// GetQuestions: return a slice of Problem type and return InvalidFileType/InvalidSchema error if file type or schema is invalid
func GetQuestions(args []string) ([]types.Problem, error) {
	fname := args[0] //file name
	oFile, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer oFile.Close()
	problems, err := validators.IsValid(oFile)
	if err != nil {
		if schemaErrors, ok := err.(*customErrors.ErrInvalidProblems); ok {
			for _, reason := range schemaErrors.InvalidQuestions {
				for qno, r := range reason {
					fmt.Printf("Disregarding question %d because %s\n", qno, r)
				}
			}
		} else {
			return nil, err
		}
	}
	if len(problems) == 0 {
		return nil, customErrors.ErrInvalidSchema
	}
	return problems, nil
}
