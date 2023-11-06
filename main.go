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

	marks := questionUser(probarr)

	fmt.Printf("You got %d/%d correct!\n", marks, len(probarr))

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
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
func questionUser(questionarr []problem) int {
	marks := 0
	var ans string

	for _, prob := range questionarr {
		fmt.Print(prob.question, ": ")
		fmt.Scan(&ans)

		if ans == prob.answer {
			marks++
		}
	}

	return marks
}
