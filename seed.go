package fuzzmath

import (
	"strings"
	"testing"
)

type seed struct {
	uints []uint64
	neg   bool
}

func (s seed) String(maxDigits int) string {
	strLen := maxDigits
	if s.neg {
		strLen++
	}

	sb := &strings.Builder{}
	sb.Grow(strLen)

	if s.neg {
		sb.WriteRune('-')
	}

	for _, digit := range s.uints {
		uintToString(sb, digit)
	}

	return sb.String()
}

func seedToStringAdapter(f *testing.F, maxDigits int, handler func(t *testing.T, numbers []string)) func(t *testing.T, seeds []seed) {
	f.Helper()

	return func(t *testing.T, seeds []seed) {
		f.Helper()
		t.Helper()

		parsedNumbers := make([]string, 0, len(seeds))
		for _, s := range seeds {
			parsedNumbers = append(parsedNumbers, s.String(maxDigits))
		}

		handler(t, parsedNumbers)
	}
}
