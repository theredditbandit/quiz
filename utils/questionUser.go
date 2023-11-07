package utils

import (
	"bufio"
	"fmt"
	"os"
	"quiz/types"
	"strings"
	"time"
)

// takes in a problem array and total time limit prints the questions, returns marks and errors
func QuestionUser(questions []types.Problem, totalTime int, reader types.ReaderFunc, testTimer types.TimerFunc) (int, error) {
	marks := 0
	attempted := 0
	var errors []types.UserError
	var qEval QuizEvaluation
	answerCh := make(chan string)
	for pid, problem := range questions {
		fmt.Printf("%d) %s = ", pid+1, problem.Question)
		go func() {
			answer, _ := reader()
			answerCh <- answer
		}()

		select {
		case <-testTimer(time.Duration(totalTime) * time.Second):
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
				errors = append(errors, types.UserError{GivenProb: problem, UserAns: ans, QuesNo: pid + 1})
			}
		}
	}

	if len(errors) > 0 {
		qEval.IncorrectlyAttempted = errors
		return marks, qEval
	}

	return marks, nil
}

func ConsoleReader() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	return strings.TrimSpace(answer), err
}

func QuizTimer(d time.Duration) <-chan time.Time {
	return time.After(d)
}
