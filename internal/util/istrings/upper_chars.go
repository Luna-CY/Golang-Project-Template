package istrings

import (
	"regexp"
	"strings"
)

func GetUpperChars(data string) string {
	return strings.Join(regexp.MustCompile(`[A-Z]`).FindAllString(data, -1), "")
}
