package fuzzdecimal

import (
	"fmt"
	"testing"
)

func seedToStringFunc(f *testing.F, cfg config, handler func(t *testing.T, strNumbers []string)) func(t *testing.T, seeds []seed) {
	f.Helper()

	return func(t *testing.T, seeds []seed) {
		t.Helper()

		parsedNumbers := make([]string, 0, len(seeds))
		for i, s := range seeds {
			parsedNumbers = append(parsedNumbers, s.string(cfg.decimals[i]))
		}

		handler(t, parsedNumbers)
	}
}

func parseStringSlice[N any](t *T, strNumbers []string, parseNumberFunc func(t *T, s string) (N, error)) ([]N, error) {
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
