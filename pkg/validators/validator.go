package validators

import (
	"os"
	"quiz/pkg/customErrors"
	"quiz/pkg/types"
	"quiz/pkg/utils"
)

// IsValid: returns true if schema of provided file matches supported csv schema
func IsValid(openFile *os.File) (bool, []types.Problem, error) {
	switch utils.GetFileType(openFile.Name()) {
	case "csv":
		return csvValidator(openFile)
	case "json":
		return jsonValidator(openFile)
	default:
		return false, nil, customErrors.ErrInvalidFileType
	}

}
