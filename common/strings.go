package common

import (
	"fmt"
	"regexp"
	"strings"
)

// NewlineRegex is a compiled regex to detect newlines
var NewlineRegex = regexp.MustCompile(`\r?\n`)

// Pluralize will return either the singular or pluralized string
func Pluralize(count int, singular string, plural string) string {

	if count < 0 {
		count = 0
	}

	if count == 1 {
		return singular
	}

	return plural
}


// TruncateString will cut a string to the length give in the function argument and return
// the cut string with an elipse
func TruncateString(str string, maxLength int) string {
	str = NewlineRegex.ReplaceAllString(str, " ")
	str = strings.TrimSpace(str)
	if len(str) > maxLength {
		str = fmt.Sprintf("%s...", str[0:maxLength])
	}
	return str
}
