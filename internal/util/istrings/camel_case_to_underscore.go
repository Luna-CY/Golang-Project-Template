package istrings

import "unicode"

func CamelCaseToUnderscore(s string) string {
	var output []rune

	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}

		if unicode.IsUpper(r) {
			output = append(output, '_')
		}

		output = append(output, unicode.ToLower(r))
	}

	return string(output)
}
