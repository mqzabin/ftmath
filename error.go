package fuzzdecimal

import (
	"fmt"
)

func errorMessageFromValues(prefix string, values []string) string {
	var errorMessage string
	for index, value := range values {
		errorMessage += fmt.Sprintf("%sindex %d value: '%s'\n", prefix, index, value)
	}

	return errorMessage
}
