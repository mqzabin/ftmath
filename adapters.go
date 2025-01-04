package fuzzdecimal

import (
	"fmt"
	"testing"
)

func seedToStringAdapter(maxDigits int, handler func(t *testing.T, strNumbers []string)) func(t *testing.T, seeds []Seed) {
	return func(t *testing.T, seeds []Seed) {
		t.Helper()

		parsedNumbers := make([]string, 0, len(seeds))
		for _, s := range seeds {
			parsedNumbers = append(parsedNumbers, s.String(maxDigits))
		}

		handler(t, parsedNumbers)
	}
}

func parseStringSlice[N any](t *testing.T, strNumbers []string, parseNumberFunc func(t *testing.T, s string) (N, error)) ([]N, error) {
	t.Helper()

	parsedNumbers := make([]N, 0, len(strNumbers))
	for index, strNumber := range strNumbers {
		parsedNumber, err := parseNumberFunc(t, strNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to parse number string '%s' at index %d: %v", strNumber, index, err)
		}

		parsedNumbers = append(parsedNumbers, parsedNumber)
	}

	return parsedNumbers, nil
}
