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

	filereader := csv.NewReader(file)
	lines, err := filereader.ReadAll()

	if err != nil {
		utils.Exit("Could not parse provided csv")
	}

	probarr := parseLines(lines) // array of problem type

	marks , err := questionUser(probarr)
	
	fmt.Printf("You got %d/%d correct!\n", marks, len(probarr))


	if err != nil {
		userErrs, _ := err.(static.UserErrors)
		userErrs.PrintErrors()
	}
}

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

// takes in a problem array and returns the users marks
func questionUser(questionarr []static.Problem) (int, error) {

	marks := 0
	reader := bufio.NewReader(os.Stdin)
	var err []static.UserErr
	var userErrors static.UserErrors

	for _, prob := range questionarr {
		fmt.Print(prob.Question, ": ")
		ans, _ := reader.ReadString('\n')
		ans = strings.TrimSpace(ans)
		if ans == prob.Answer {
			marks++
		} else {
			err = append(err, static.UserErr{GivenProb: prob, UserAns: ans})
		}
	}

	if len(err) > 0 {
		userErrors.Code = len(err)
		userErrors.Msg = "You have Mistakes" 
		userErrors.Errors = err
		return marks , userErrors
	}

	return marks , nil
}
