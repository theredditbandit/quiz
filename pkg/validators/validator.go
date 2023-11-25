package validators

import (
	"os"
	"quiz/pkg/customErrors"
	"quiz/pkg/types"
	"quiz/pkg/utils"
)

// IsValid: validates and returns problems slice and schema errors if any
func IsValid(openFile *os.File) ([]types.Problem, error) {
	switch utils.GetFileType(openFile.Name()) {
	case "csv":
		return csvValidator(openFile)
	case "json":
		return jsonValidator(openFile)
	default:
		return nil, customErrors.ErrInvalidFileType
	}

}
