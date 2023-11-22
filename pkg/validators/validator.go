package validators

import (
	"quiz/pkg/customErrors"
	"quiz/pkg/utils"
)

// IsValid: returns true if schema of provided file matches supported csv schema
func IsValid(file string) (bool, error) {
	switch utils.GetFileType(file) {
	case "csv":
		return csvValidator(file)
	case "json":
		return jsonValidator(file)
	default:
		return false, customErrors.ErrInvalidFileType
	}

}
