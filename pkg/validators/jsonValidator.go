package validators

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"quiz/pkg/types"
)

func jsonValidator(jsonFile string) bool {
	var problems []types.Problem

	oFile, err := os.Open(jsonFile)
	if err != nil {
		return false
	}
	data, err := io.ReadAll(oFile)
	if err != nil {
		return false
	}
	err = json.Unmarshal(data, &problems)

	fmt.Printf("Data read :\n%+v\n", problems[1])
	return err == nil

}
