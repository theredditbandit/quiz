package parsers

import "quiz/pkg/customTypes"

// takes a 2d array of type [ [q1 ,a1], [q2 ,a2] . . . ] and
// returns an array of type static.Problem
func LinesToProblems(lines [][]string) []customTypes.Problem {
	ret := make([]customTypes.Problem, len(lines))
	question := 0
	answer := 1
	for idx, qa := range lines {
		ret[idx].Question = qa[question]
		ret[idx].Answer = qa[answer]

	}
	return ret
}
