package utils

import (
	"quiz/types"
	"testing"
	"time"
)

func TestQuestionUser(t *testing.T) {
	type args struct {
		questions []types.Problem
		totalTime int
		reader    types.ReaderFunc
		testTimer types.TimerFunc
	}
	tests := []struct {
		name    string
		args    args
		marks    int
		wantErr bool
	}{
		{
			name: "Test with valid input",
			args: args{
				questions: []types.Problem{
					{
						Question: "2+2?",
						Answer:   "4",
					},
				},
				totalTime: 10,
				reader: func() (string, error) {
					return "4", nil
				},
				testTimer: func(d time.Duration) <-chan time.Time {
					return time.After(1 * time.Second)
				},
			},
			marks:    1,
			wantErr: false,
		},
		{
			name: "Test with invalid input",
			args: args{
				questions: []types.Problem{
					{
						Question: "2+2",
						Answer:   "4",
					},
				},
				totalTime: 10,
				reader: func() (string, error) {
					return "5", nil
				},
				testTimer: func(d time.Duration) <-chan time.Time {
					return time.After(1 * time.Second)
				},
			},
			marks:    0,
			wantErr: true,
		},
		{
			name: "Test with timeout",
			args: args{
				questions: []types.Problem{
					{
						Question: "2+2",
						Answer:   "4",
					},
				},
				totalTime: 1,
				reader: func() (string, error) {
					return "", nil
				},
				testTimer: func(d time.Duration) <-chan time.Time {
					return time.After(2 * time.Second)
				},
			},
			marks:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QuestionUser(tt.args.questions, tt.args.totalTime, tt.args.reader, tt.args.testTimer)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuestionUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.marks {
				t.Errorf("QuestionUser() = %v, want %v", got, tt.marks)
			}
		})
	}
}