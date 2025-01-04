package fuzzdecimal

import (
	"fmt"
)

func errorMessageFromValues(values []string) string {
	errorMessage := "\n"
	for index, value := range values {
		errorMessage += fmt.Sprintf("index %d value: '%s'\n", index, value)
	}

	return errorMessage
}
