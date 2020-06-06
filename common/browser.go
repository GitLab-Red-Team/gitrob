package common

import (
	"fmt"
	"strings"
)

// UserAgent will store the proper agent
var UserAgent = fmt.Sprintf("%s v%s", Name, Version)

// CleanUrlSpaces will replaces spaces in a string with dashes so that they can be used in urls
func CleanUrlSpaces(dirtyStrings ...string) []string {
	var result []string
	for _, s := range dirtyStrings {
		result = append(result, strings.ReplaceAll(s, " ", "-"))
	}
	return result
}
