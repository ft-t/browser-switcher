package escaper

import (
	"fmt"
	"strings"
)

var charsToEscape = []string{
	"&",
}

func Escape(
	input string,
) string {
	for _, char := range charsToEscape {
		input = strings.ReplaceAll(input, char, fmt.Sprintf(`"%s"`, char))
	}

	return input
}
