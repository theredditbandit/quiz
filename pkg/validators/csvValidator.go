package validators

import (
	"encoding/csv"
	"fmt"
	"os"
	"quiz/pkg/customErrors"
	"quiz/pkg/types"
)

// csvValidator: validates the CSV file and returns a slice of the problems
func csvValidator(oFile *os.File) ([]types.Problem, error) {
	reader := csv.NewReader(oFile)
	firstLine, err := reader.Read()
	if err != nil {
		return nil, err
	}
	if len(firstLine) != 2 {
		return nil, customErrors.ErrInvalidSchema
	}
	fmt.Println("CSV Schema is valid") // [ ] TODO  log.debug this in when logging is implemented
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	problems := make([]types.Problem, len(lines))
	for idx, line := range lines {
		problems[idx].Question = line[0] // line[0] is the question
		problems[idx].Answer = line[1]   // line[1] is the answer
	}
	return problems, nil
}
