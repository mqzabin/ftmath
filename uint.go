package fuzzdecimal

const (
	// maxDigitsPerUint is the number of digits that can be safely represented by an uint64.
	// 18446744073709551615 is the maximum value that can be represented by an uint64,
	// and the max number that could be completely represented by it, is:
	// 9999999999999999999, which has 19 digits.
	maxDigitsPerUint = 19

	maxUintValue = 9999999999999999999

	base = 10
)

var digits = []string{
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

var pow10 = []uint64{
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

func normalizeUint(n uint64, maxDigits int) uint64 {
	if maxDigits > maxDigitsPerUint {
		panic("got maxDigits > maxDigitsPerUint")
	}

	return n % pow10[maxDigits]
}

func uintToString(n uint64) string {
	var result string

	for index := maxDigitsPerUint - 1; index >= 0; index-- {
		result += digits[(n/pow10[index])%base]

		n %= pow10[index]
	}

	return result
}

func uintsPerNumber(maxDigits int) int {
	if maxDigits == 0 {
		return 0
	}

	return ((maxDigits - 1) / maxDigitsPerUint) + 1
}
