package fdlib

import (
	"fmt"
	"testing"
)

// SeedsFuncToStringsFunc will convert a function that takes a slice of Seed and to a function that takes a slice of decimal strings.
func SeedsFuncToStringsFunc(f *testing.F, cfg Config, handler func(t *testing.T, strNumbers []string)) func(t *testing.T, seeds []Seed) {
	f.Helper()

	return func(t *testing.T, seeds []Seed) {
		t.Helper()

		parsedNumbers := make([]string, 0, len(seeds))
		for i, s := range seeds {
			parsedNumbers = append(parsedNumbers, s.String(t, cfg.Decimals[i]))
		}

		handler(t, parsedNumbers)
	}
}

// ParseStringSliceToDecimalSlice will parse a slice of decimal strings to a slice of Decimals types.
func ParseStringSliceToDecimalSlice[Decimal any](t *testing.T, decimalStrings []string, parseDecimalFunc func(t *testing.T, s string) (Decimal, error)) ([]Decimal, error) {
	parsedNumbers := make([]Decimal, 0, len(decimalStrings))
	for index, decimalString := range decimalStrings {
		parsedNumber, err := parseDecimalFunc(t, decimalString)
		if err != nil {
			return nil, fmt.Errorf("failed to parse decimal string '%s' at index %d: %v", decimalString, index, err)
		}

		parsedNumbers = append(parsedNumbers, parsedNumber)
	}

	return parsedNumbers, nil
}
