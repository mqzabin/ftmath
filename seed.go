package ftmath

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

func seedToStringAdapter[N any](f *testing.F, maxDigits int, parseFunc func(t *testing.T, s string) N, handler func(t *testing.T, numbers []N)) func(t *testing.T, seeds []seed) {
	f.Helper()

	return func(t *testing.T, seeds []seed) {
		f.Helper()
		t.Helper()

		parsedNumbers := make([]N, 0, len(seeds))
		for _, s := range seeds {
			parsedNumbers = append(parsedNumbers, parseFunc(t, s.String(maxDigits)))
		}

		handler(t, parsedNumbers)
	}
}
