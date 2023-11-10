package validators

import (
	"encoding/csv"
	"os"
	"quiz/pkg/utils"
)

var CSVSchema = []string{"Qno ,Question,Answer"}

func csvValidator(csvFile string) bool {
	openFile, err := os.Open(csvFile)
	if err != nil {
		utils.ExitWithMessage("Something went wrong while opening the file , could not validate.", 1)
	}
	reader := csv.NewReader(openFile)
	firstLine, err := reader.Read()
	if err != nil {
		utils.ExitWithMessage("Something went wrong while reading the CSV , could not validate", 1)
	}

	if len(firstLine) == 2 {
		return true
	} else {
		return false
	}
}
