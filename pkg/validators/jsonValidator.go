package validators

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"quiz/pkg/customErrors"
	"quiz/pkg/types"
	"strings"
)

// jsonValidator: returns true if at least one of the questions supplied has a valid schema , false otherwise. error contains the first validation error encountered.
func jsonValidator(oFile *os.File) ([]types.Problem, error) {
	problems, err := getJsonData(oFile)
	validProblems := make([]types.Problem, 0)
	if err != nil {
		return nil, err
	}
	var errInvalidProblems customErrors.ErrInvalidProblems
	for _, p := range problems {
		reason, isValid := validate(p)
		if !isValid {
			errInvalidProblems.InvalidQuestions = append(errInvalidProblems.InvalidQuestions, reason)
		} else {
			fmt.Printf("Q%d is valid \n", p.QuestionNumber)
			validProblems = append(validProblems, p)
		}
	}
	return validProblems, &errInvalidProblems
}

// validate: validates a single problem and returns reason along with validity
func validate(p types.Problem) (map[int]string, bool) {
	reason := make(map[int]string)
	if p.QuestionNumber <= 0 {
		reason[p.QuestionNumber] = "Question number cannot be negative or zero"
		return reason, false
	}
	if strings.TrimSpace(p.Question) == "" {
		reason[p.QuestionNumber] = "Question text cannot be empty"
		return reason, false
	}
	if p.IsMCQTypeQuestion && len(strings.TrimSpace(p.Answer)) > 0 {
		reason[p.QuestionNumber] = "MCQ type question cannot have answer specified, use Options instead"
		return reason, false
	} else if !p.IsMCQTypeQuestion {
		if strings.TrimSpace(p.Answer) == "" {
			reason[p.QuestionNumber] = "Non-MCQ type question must have answer specified"
			return reason, false
		} else if p.AllowMultipleAns {
			reason[p.QuestionNumber] = "Non-MCQ type question cannot allow multiple answers"
			return reason, false
		} else if len(p.Options) > 0 {
			reason[p.QuestionNumber] = "Non-MCQ type question cannot have options specified"
			return reason, false
		} else if len(p.MCQAnswers) > 0 {
			reason[p.QuestionNumber] = "Non-MCQ type question cannot have MCQ answers specified"
			return reason, false
		}
	}
	if p.AllowMultipleAns && len(p.MCQAnswers) == 1 {
		reason[p.QuestionNumber] = "Question cannot allow multiple answers but have only 1 MCQ answer specified"
		return reason, false
	}
	if len(p.Options) == 1 {
		reason[p.QuestionNumber] = "MCQ type question cannot have only 1 option specified"
		return reason, false
	}
	if len(p.MCQAnswers) > 0 {
		for _, ans := range p.MCQAnswers {
			if _, ok := p.Options[ans]; !ok {
				reason[p.QuestionNumber] = fmt.Sprintf("MCQ option %s is not present in options", ans)
				return reason, false
			}
		}
	}
	if p.IsTimed {
		if p.Time.Time <= 0 {
			reason[p.QuestionNumber] = "Time value must be positive for timed questions"
			return reason, false
		}
		if p.Time.Unit == "" {
			reason[p.QuestionNumber] = "Defaulting to seconds for timed questions without unit specified"
			return reason, true
		} else if p.Time.Unit != "hr" && p.Time.Unit != "min" && p.Time.Unit != "sec" {
			reason[p.QuestionNumber] = "Time unit must be 'hr', 'min' or 'sec' for timed questions"
			return reason, false
		}
	}

	return nil, true
}

// getJsonData: Unmarshall the json in the given file and return a slice of type Problem
func getJsonData(oFile *os.File) ([]types.Problem, error) {
	var problems []types.Problem
	data, err := io.ReadAll(oFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &problems)
	if err != nil {
		return nil, err
	}
	return problems, nil
}
