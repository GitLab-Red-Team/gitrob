package common

import "os"

// FileExists will check if a file exists and return a bool based on the result
func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
