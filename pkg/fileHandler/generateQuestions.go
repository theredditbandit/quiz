package fileHandler

import (
	"encoding/json"
	"os"
	"quiz/pkg/types"
)

// GenBoilerplate function generates boilerplate JSON file for MCQ type questions
func GenBoilerplate(n int, name string) error { // [ ] TODO: add a spinner maybe?
	schemas := make([]types.Problem, n)

	for i := range schemas {
		schemas[i].QuestionNumber = i + 1
		schemas[i].AllowMultipleAns = false
		schemas[i].Question = "Question goes here"
		schemas[i].Answer = "Ignore this field since the question is MCQ"
		schemas[i].Options = map[string]string{"a": "change this , add more options as needed", "b": "change this , add more options as needed"}
		schemas[i].MCQAnswers = []string{"a"}
	}

	jsonData, err := json.MarshalIndent(schemas, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(name+".json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
