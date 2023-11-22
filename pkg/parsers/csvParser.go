package parsers

import "quiz/pkg/types"

// LinesToProblems: takes a 2d array of type [ [q1 ,a1], [q2 ,a2] . . . ] and
// returns an array of problem type
func LinesToProblems(lines [][]string) []types.Problem {
	ret := make([]types.Problem, len(lines))
	question := 0
	answer := 1
	for idx, qa := range lines {
		ret[idx].Question = qa[question]
		ret[idx].Answer = qa[answer]

	}
	return ret
}
