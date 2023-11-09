package utils

import (
	"flag"
)

func GetArguments() (*string, *int, *bool) {
	problemFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	testTime := flag.Int("time", 30, "Specifies the total time the quiz is going to run for.")
	shuffle := flag.Bool("shuffle", false, "Whether or not to shuffle questions while asking.")
	flag.Parse()
	return problemFile, testTime, shuffle
}
