package escaper

import (
	"fmt"
	"strings"
)

func Unescape(
	input string,
) string {
	for _, char := range charsToEscape {
		input = strings.ReplaceAll(input, fmt.Sprintf(`"%s"`, char), char)
	}

	return input
}
