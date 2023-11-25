package customErrors

import (
	"errors"
	"fmt"
)

var ErrInvalidSchema = errors.New("invalidSchema")
var ErrInvalidFileType = errors.New("invalidFileType")

type ErrInvalidProblems struct {
	InvalidQuestions []map[int]string
}

func (e *ErrInvalidProblems) Error() string {
	return fmt.Sprintf("%d problems have invalid schema", len(e.InvalidQuestions))
}
