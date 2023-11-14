package testUser

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"quiz/pkg/types"
	"strings"
	"time"
)

// Questionuser: takes in a problem array and total time limit prints the questions, returns marks and errors
func QuestionUser(questions []types.Problem, totalTime types.TimeConf, reader types.ReaderFunc, testTimer types.TimerFunc) (int, error) {
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
		case <-testTimer(totalTime):
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

func QuizTimer(tconf types.TimeConf) <-chan time.Time {
	qtime := tconf.Time
	var t time.Duration
	if qtime > 0 {
		switch tconf.Unit {
		case "sec":
			t = time.Duration(qtime) * time.Second
		case "min":
			t = time.Duration(qtime) * time.Minute
		case "hour":
			t = time.Duration(qtime) * time.Hour
		}

	} else {
		t = time.Duration(math.MaxInt64)
	}

	return time.After(t)
}
