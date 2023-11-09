package parsers

import (
	"quiz/pkg/customTypes"
	"reflect"
	"testing"
)

func TestParseLines(t *testing.T) {
	type args struct {
		lines [][]string
	}
	tests := []struct {
		name string
		args args
		want []customTypes.Problem
	}{
		{
			name: "empty input",
			args: args{lines: [][]string{}},
			want: []customTypes.Problem{},
		},
		{
			name: "single input",
			args: args{lines: [][]string{{"question1", "answer1"}}},
			want: []customTypes.Problem{{Question: "question1", Answer: "answer1"}},
		},
		{
			name: "multiple inputs",
			args: args{lines: [][]string{{"question1", "answer1"}, {"question2", "answer2"}}},
			want: []customTypes.Problem{{Question: "question1", Answer: "answer1"}, {Question: "question2", Answer: "answer2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinesToProblems(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
