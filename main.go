package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"quiz/static"
	"quiz/utils"
	"strings"
	"time"
)

func main() {
	problemFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	testTime := flag.Int("time", 30, "Specifies the total time the quiz is going to run for.")
	flag.Parse()
	file, err := os.Open(*problemFile)

	if err != nil {
		utils.Exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *problemFile))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		utils.Exit("Could not parse provided csv")
	}

	problems := parseLines(lines) // array of problem type

	marks, err := questionUser(problems, *testTime)

	fmt.Printf("You got %d/%d correct!\n", marks, len(problems))

	if err != nil {
		userErrs, _ := err.(static.QuizEvaluation)
		userErrs.PrintErrors()

		if userErrs.Unattempted {
			userErrs.PrintUnattempted()
		}
	}
}

// takes a 2d array of type [ [q1 ,a1], [q2 ,a2] . . . ] and
// returns an array of type static.Problem
func parseLines(lines [][]string) []static.Problem {
	ret := make([]static.Problem, len(lines))
	question := 0
	answer := 1
	for idx, qa := range lines {
		ret[idx].Question = qa[question]
		ret[idx].Answer = qa[answer]

	}
	return ret
}

// takes in a problem array and total time limit prints the questions, returns marks and errors
func questionUser(questions []static.Problem, totalTime int) (int, error) {
	marks := 0
	attempted := 0
	reader := bufio.NewReader(os.Stdin)
	var errors []static.UserError
	var qEval static.QuizEvaluation
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
