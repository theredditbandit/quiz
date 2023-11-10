package utils

import "strings"

func GetFileType(fileName string) (fileType string) {
	defer func() {
		if r := recover(); r != nil {
			fileType = "invalid"
		}
	}()
	fsplit := strings.Split(fileName, ".")
	return fsplit[1]
}
