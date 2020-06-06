package common

import (
	"fmt"
	"os"
)

// FileExists will check if a file exists and return a bool based on the result
func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err != nil {
			fmt.Println("The file could not be checked")
			// TODO Handle the error msg
		}
		return false
	}

	return true
}
