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
)

func main() {
	problemFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
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

	marks, err := questionUser(problems)

	fmt.Printf("You got %d/%d correct!\n", marks, len(problems))

	if err != nil {
		userErrs, _ := err.(static.QuizErrors)
		userErrs.PrintErrors()
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

// takes in a problem array and prints the questions, returns marks and errors
func questionUser(questions []static.Problem) (int, error) {

	marks := 0
	reader := bufio.NewReader(os.Stdin)
	var errors []static.UserError
	var qErrors static.QuizErrors

	for _, problem := range questions {
		fmt.Print(problem.Question, ": ")
		ans, _ := reader.ReadString('\n')
		ans = strings.TrimSpace(ans)
		if ans == problem.Answer {
			marks++
		} else {
			errors = append(errors, static.UserError{GivenProb: problem, UserAns: ans})
		}
	}

	if len(errors) > 0 {
		qErrors.Code = len(errors)
		qErrors.Msg = "You have Mistakes"
		qErrors.Errors = errors
		return marks, qErrors
	}

	return marks, nil
}
