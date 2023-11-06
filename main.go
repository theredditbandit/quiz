package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	problemFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	file, err := os.Open(*problemFile)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *problemFile))
	}

	filereader := csv.NewReader(file)
	lines, err := filereader.ReadAll()

	if err != nil {
		exit("Could not parse provided csv")
	}

	probarr := parseLines(lines) // array of problem type

	sol := questionUser(probarr)

	fmt.Printf("You got %d/%d correct!\n", sol.marks, len(probarr))

	if len(sol.err) != 0 {
		for _, val := range sol.err {
			fmt.Println("Question ", val.question.question, "expected answer ", val.question.answer, "instead got ", val.useranswer)
		}

	}

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}

type userSol struct {
	marks int
	err   []usererr
}

type usererr struct {
	question   problem
	useranswer string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	question := 0
	answer := 1
	for idx, qa := range lines {
		ret[idx].question = qa[question]
		ret[idx].answer = qa[answer]

	}
	return ret
}

// takes in a problem array and returns the users marks
func questionUser(questionarr []problem) userSol {
	var sol userSol
	sol.marks = 0

	var ans string

	for _, prob := range questionarr {
		fmt.Print(prob.question, ": ")
		fmt.Scan(&ans)
		if ans == prob.answer {
			sol.marks++
		} else {
			sol.err = append(sol.err, usererr{question: prob, useranswer: ans})
		}
	}

	return sol
}
