package codesafe

import (
	"strings"
	"unicode"
)

func ToCamelCase(input string) string {
	// Split the input string based on underscores, hyphens, or spaces
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == '_' || r == '-' || unicode.IsSpace(r)
	})

	// Capitalize the first letter of each word (except the first one)
	for i := 0; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	// Concatenate the words to form CamelCase
	result := strings.Join(words, "")

	return result
}
