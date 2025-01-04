package fuzzdecimal

import (
	"strings"
)

const (
	// uintSafeDigits is the number of digits that can be safely represented by an uint64.
	// 18446744073709551615 is the maximum value that can be represented by an uint64,
	// and the max number that could be completely represented by it, is:
	// 9999999999999999999, which has 19 digits.
	uintSafeDigits = 19

	base = 10
)

func uintToString(sb *strings.Builder, n uint64) {
	var digitRune rune

	sbCap := sb.Cap()
	sbLen := sb.Len()

	for range uintSafeDigits {
		if sbLen == sbCap {
			return
		}

		switch n % base {
		case 0:
			digitRune = '0'
		case 1:
			digitRune = '1'
		case 2:
			digitRune = '2'
		case 3:
			digitRune = '3'
		case 4:
			digitRune = '4'
		case 5:
			digitRune = '5'
		case 6:
			digitRune = '6'
		case 7:
			digitRune = '7'
		case 8:
			digitRune = '8'
		case 9:
			digitRune = '9'
		}

		_, _ = sb.WriteRune(digitRune)

		sbLen++
		n /= base
	}
}

func UintsPerNumber(maxDigits int) int {
	return maxDigits / uintSafeDigits
}
