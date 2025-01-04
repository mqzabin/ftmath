package fuzzdecimal

import (
	"strings"
)

type Seed struct {
	uints []uint64
	neg   bool
}

func (s Seed) String(maxDigits int) string {
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
