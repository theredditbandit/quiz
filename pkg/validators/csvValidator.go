package validators

import (
	"encoding/csv"
	"os"
	"quiz/pkg/utils"
)

func csvValidator(csvFile string) bool {
	oFile, err := os.Open(csvFile)
	if err != nil {
		utils.ExitWithMessage("Something went wrong while opening the file , could not validate.", 1)
	}
	defer oFile.Close()
	reader := csv.NewReader(oFile)
	firstLine, err := reader.Read()
	if err != nil {
		utils.ExitWithMessage("Something went wrong while reading the CSV , could not validate", 1)
	}

	return len(firstLine) == 2 // if the first line has 2 columns then it is a valid csv
}
