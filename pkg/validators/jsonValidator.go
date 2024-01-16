package validators

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"quiz/pkg/colors"
	"quiz/pkg/customErrors"
	"quiz/pkg/types"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh/spinner"
)

// jsonValidator: returns true if at least one of the questions supplied has a valid schema , false otherwise. error contains the first validation error encountered.
func jsonValidator(oFile *os.File) ([]types.Problem, error) {
	probChan := make(chan []types.Problem, 2)
	errChan := make(chan error, 3)
	validationSummary := make(chan string, 2)
	problems, err := getJsonData(oFile)
	validator := func() { // NOTE : Maybe move this function to another file ?
		time.Sleep(1 * time.Second) // PERF : sleeping because I worked hard to implement the spinner and I want to see it
		validProblems := make([]types.Problem, 0)
		if err != nil {
			probChan <- nil
			errChan <- err
		}
		var errorsAndWarnings customErrors.ErrInvalidProblems
		for _, p := range problems {
			reason, isValid := validateOne(p)
			if !isValid {
				errorsAndWarnings.InvalidQuestions = append(errorsAndWarnings.InvalidQuestions, reason)
			} else {
				if len(reason) != 0 { // we only get here when the question is valid but a reason is still returned
					errorsAndWarnings.Warnings = append(errorsAndWarnings.Warnings, reason)
				}
				validProblems = append(validProblems, p)
			}
		}
		correctQuestions := strconv.Itoa(len(validProblems))
		incorrectQuestions := strconv.Itoa(len(errorsAndWarnings.InvalidQuestions))
		warnings := strconv.Itoa(len(errorsAndWarnings.Warnings))
		if incorrectQuestions == "0" && warnings == "0" {
			validationSummary <- colors.GreatSuccess.Render("All " + correctQuestions + " questions are valid ✅")
			errChan <- nil // as there are no validation errors or warnings
		} else {
			validationSummary <- fmt.Sprintf("There's %v , %v and %v",
				colors.GreatSuccess.Render(correctQuestions+" Correct Question(s)"),
				colors.GraveError.Render(incorrectQuestions+" Incorrect Question(s)"),
				colors.MuchWarning.Render(warnings+" Warning(s)"))
		}
		probChan <- validProblems
		errChan <- &errorsAndWarnings
	}
	_ = spinner.New().Title("Validating questions . . .").Action(validator).Run()
	fmt.Println(<-validationSummary)
	return <-probChan, <-errChan
}

// validateOne validates a single problem and returns reason along with validity
// sqipcq: GO-R1005
func validateOne(p types.Problem) (map[int]string, bool) {
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
		reason, isNonMCQValid := validateNonMCQ(p)
		if !isNonMCQValid {
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
			return reason, true // warning
		} else if p.Time.Unit != "hr" && p.Time.Unit != "min" && p.Time.Unit != "sec" {
			reason[p.QuestionNumber] = "Time unit must be 'hr', 'min' or 'sec' for timed questions"
			return reason, false
		}
	}

	if p.DisplayExplanation && (strings.TrimSpace(p.Explanation) == "" && strings.TrimSpace(p.Reference) == "") {
		reason[p.QuestionNumber] = "Explanation field cannot be empty if it's to be displayed"
		return reason, false
	}
	if !p.DisplayExplanation && len(strings.TrimSpace(p.Explanation)) > 0 {
		reason[p.QuestionNumber] = "Ignoring Explanation: DisplayExplanation set to false but Explanation provided"
		return reason, true // warning
	}
	if !p.DisplayExplanation && len(strings.TrimSpace(p.Reference)) > 0 {
		reason[p.QuestionNumber] = "Ignoring Reference: DisplayExplanation set to false but Reference(s) provided"
		return reason, true // warning
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

// validateNonMCQ validates a non-MCQ type question and returns reason along with validity
func validateNonMCQ(p types.Problem) (map[int]string, bool) {
	reason := make(map[int]string)
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
	return nil, true
}
