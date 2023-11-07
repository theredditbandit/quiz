package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"quiz/utils"
	"quiz/utils/parsers"
)

func main() {
	problemFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	testTime := flag.Int("time", 30, "Specifies the total time the quiz is going to run for.")
	shuffle := flag.Bool("shuffle", false, "Whether or not to shuffle questions while asking.")
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
	problems := parsers.ParseLines(lines) // array of problem type

	if *shuffle {
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	marks, err := utils.QuestionUser(problems, *testTime)
	fmt.Printf("You got %d/%d correct!\n", marks, len(problems))

	if err != nil {
		userErrs, _ := err.(utils.QuizEvaluation)
		userErrs.PrintErrors()

		if userErrs.Unattempted {
			userErrs.PrintUnattempted()
		}
	}
}