package utils

import (
	"fmt"
	"quiz/static"
	"time"
	"bufio"
	"strings"
	"os"
)

// takes in a problem array and total time limit prints the questions, returns marks and errors
func QuestionUser(questions []static.Problem, totalTime int) (int, error) {
	marks := 0
	attempted := 0
	reader := bufio.NewReader(os.Stdin)
	var errors []static.UserError
	var qEval QuizEvaluation
	answerCh := make(chan string)
	testTimer := time.NewTimer(time.Duration(totalTime) * time.Second)

	for pid, problem := range questions {
		fmt.Print(pid+1, ")  ", problem.Question, " = ")
		go func() {
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(answer)
			answerCh <- answer
		}()

		select {
		case <-testTimer.C:
			fmt.Println("\nTime limit reached!")
			if attempted != len(questions) {
				// timer ran out , questions missed
				qEval.UnattemptedQuestions = questions[attempted:]
				qEval.Unattempted = true
				qEval.Attempted = attempted
			}
			if len(errors) > 0 {
				qEval.IncorrectlyAttempted = errors
			}
			return marks, qEval

		case ans := <-answerCh:
			attempted++
			if ans == strings.TrimSpace(problem.Answer) {
				marks++
			} else {
				errors = append(errors, static.UserError{GivenProb: problem, UserAns: ans, QuesNo: pid + 1})
			}
		}
	}

	if len(errors) > 0 {
		qEval.IncorrectlyAttempted = errors
		return marks, qEval
	}

	return marks, nil
}
