package examples

import (
	"encoding/json"
	"os"
	"quiz/pkg/types"
)

// GenExamples generates a json file of 5 valid questions
func GenExamples() error {
	var exampleQuestions []types.Problem

	var q1 types.Problem
	q1.QuestionNumber = 1
	q1.Question = "What's 2+2"
	q1.IsMCQTypeQuestion = false
	q1.Answer = "4"
	q1.AllowMultipleAns = false
	q1.IsTimed = true
	q1.Time.Time = 10
	q1.Time.Unit = "sec"
	q1.MarksIfCorrect = 4
	q1.MarksIfIncorrect = -1
	q1.Skippable = false
	exampleQuestions = append(exampleQuestions, q1)

	var q2 types.Problem
	q2.QuestionNumber = 2
	q2.Question = "What's 2+3*(4-6)"
	q2.IsMCQTypeQuestion = true
	q2.AllowMultipleAns = false
	q2.Options = map[string]string{"a": "-10", "b": "-4", "c": "4", "d": "10"}
	q2.MCQAnswers = []string{"a"}
	q2.IsTimed = false
	q2.MarksIfCorrect = 4
	q2.MarksIfIncorrect = -1
	q2.Skippable = true
	exampleQuestions = append(exampleQuestions, q2)

	var q3 types.Problem
	q3.QuestionNumber = 3
	q3.Question = "What's the square root of 4"
	q3.IsMCQTypeQuestion = true
	q3.AllowMultipleAns = true
	q3.Options = map[string]string{"a": "2", "b": "-2", "c": "4", "d": "-4"}
	q3.MCQAnswers = []string{"a", "b"}
	q3.IsTimed = true
	q3.Time.Time = 10
	q3.Time.Unit = "sec"
	q3.Skippable = false
	exampleQuestions = append(exampleQuestions, q3)

	jsonData, err := json.MarshalIndent(exampleQuestions, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile("examples.json", jsonData, 0644)

	if err != nil {
		return err
	}
	return nil
}
