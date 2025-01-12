package fdlib

import (
	"testing"
)

const (
	negativeSignS = "-"
	decimalPointS = "."
	decimalPointR = '.'
)

type Seed struct {
	// Uints stores all the uint64 used to represent the number.
	// The 0 index is the most significant digit.
	Uints []uint64
	// Neg stores the number's sign.
	Neg bool
}

// IsZero returns true if the seed represents the number 0.
func (s Seed) IsZero(t *testing.T) bool {
	t.Helper()

	if len(s.Uints) == 0 {
		return true
	}

	for _, digit := range s.Uints {
		if digit != 0 {
			return false
		}
	}

	return true
}

// String returns the string representation of the seed, according to the provided DecimalConfig.
func (s Seed) String(t *testing.T, cfg DecimalConfig) string {
	t.Helper()

	if s.IsZero(t) {
		return "0"
	}

	var str string

	for _, nUint := range s.Uints {
		str += UintToString(nUint)
	}

	maxStrLen := UintsPerNumber(cfg.MaxSignificantDigits) * MaxDigitsPerUint

	// Only write decimal point if it is within the number's length.
	if cfg.MaxDecimalPlaces > 0 && cfg.MaxDecimalPlaces <= cfg.MaxSignificantDigits {
		decimalPointIndex := maxStrLen - cfg.MaxDecimalPlaces
		str = str[:decimalPointIndex] + decimalPointS + str[decimalPointIndex:]
	}

	str = trimInsignificantDigits(t, str)

	if s.Neg {
		str = negativeSignS + str
	}

	return str
}

func trimInsignificantDigits(t *testing.T, str string) string {
	t.Helper()

	if len(str) == 0 {
		return str
	}

	firstNonZeroIndex := -1
	for i := range str {
		if str[i] != '0' {
			firstNonZeroIndex = i

			break
		}
	}

	lastNonZeroIndex := -1
	for i := len(str) - 1; i >= firstNonZeroIndex; i-- {
		if str[i] != '0' {
			lastNonZeroIndex = i

			break
		}
	}

	if firstNonZeroIndex == -1 || lastNonZeroIndex == -1 {
		return "0"
	}

	str = str[firstNonZeroIndex : lastNonZeroIndex+1]

	if str[0] == decimalPointR {
		str = "0" + str
	}

	if str[len(str)-1] == decimalPointR {
		str = str[:len(str)-1]
	}

	return str
}
