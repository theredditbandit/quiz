package validators

import (
	"encoding/csv"
	"os"
)

func csvValidator(csvFile string) (bool, error) {
	oFile, err := os.Open(csvFile)
	if err != nil {
		return false, err
	}
	defer oFile.Close()
	reader := csv.NewReader(oFile)
	firstLine, err := reader.Read()
	if err != nil {
		return false, err
	}

	return len(firstLine) == 2, nil // if the first line has 2 columns then it is a valid csv
}
