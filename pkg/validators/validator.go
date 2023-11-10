package validators

import "quiz/pkg/utils"

// returns true if schema of provided file matches supported csv schema
func IsValid(file string) bool {
	switch utils.GetFileType(file) {
	case "csv":
		return csvValidator(file)
	case "json":
		return jsonValidator(file)
	default:
		return false
	}

}
