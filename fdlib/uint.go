package fdlib

const (
	// MaxDigitsPerUint is the number of digits that can be safely represented by an uint64.
	// 18446744073709551615 is the maximum value that can be represented by an uint64,
	// and the max number that could be completely represented by it, is:
	// 9999999999999999999, which has 19 digits.
	MaxDigitsPerUint = 19
	// MaxUintValue max uint64 value that could be used by this package, according to MaxDigitsPerUint.
	MaxUintValue = 9999999999999999999
	// base is the base used to represent numbers.
	base = 10
)

var (
	// digits is a slice of string representations of numbers from 0 to 9.
	digits = []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
	// pow10 is a slice of powers of 10.
	pow10 = []uint64{
		1,
		10,
		100,
		1000,
		10000,
		100000,
		1000000,
		10000000,
		100000000,
		1000000000,
		10000000000,
		100000000000,
		1000000000000,
		10000000000000,
		100000000000000,
		1000000000000000,
		10000000000000000,
		100000000000000000,
		1000000000000000000,
		10000000000000000000,
	}
)

// NormalizeUint truncates the uint64 digits to the maxDigits.
func NormalizeUint(th TestHelper, n uint64, maxDigits int) uint64 {
	if maxDigits > MaxDigitsPerUint {
		th.Fatalf("got maxDigits %d > MaxDigitsPerUint %d", maxDigits, MaxDigitsPerUint)
	}

	return n % pow10[maxDigits]
}

// UintToString converts an uint64 to a string, with leading zeroes up to MaxDigitsPerUint.
func UintToString(n uint64) string {
	var result string

	for index := MaxDigitsPerUint - 1; index >= 0; index-- {
		result += digits[(n/pow10[index])%base]

		n %= pow10[index]
	}

	return result
}

// UintsPerNumber returns the number of uints needed to represent a number with maxSignificantDigits.
func UintsPerNumber(maxSignificantDigits int) int {
	if maxSignificantDigits == 0 {
		return 0
	}

	return ((maxSignificantDigits - 1) / MaxDigitsPerUint) + 1
}
