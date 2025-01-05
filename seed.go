package fuzzdecimal

type seed struct {
	// uints stores all the uint64 used to represent the number.
	// The 0 index is the most significant digit.
	uints []uint64
	// neg stores the number's sign.
	neg bool
}

func (s seed) isZero() bool {
	if len(s.uints) == 0 {
		return true
	}

	for _, digit := range s.uints {
		if digit != 0 {
			return false
		}
	}

	return true
}

func (s seed) string(cfg decimalConfig) string {
	if s.isZero() {
		return "0"
	}

	var str string

	for _, nUint := range s.uints {
		str += uintToString(nUint)
	}

	maxStrLen := uintsPerNumber(cfg.maxSignificantDigits) * maxDigitsPerUint

	// Only write decimal point if it is within the number's length.
	if cfg.decimalPointPosition > 0 && cfg.decimalPointPosition <= cfg.maxSignificantDigits {
		decimalPointIndex := maxStrLen - cfg.decimalPointPosition
		str = str[:decimalPointIndex] + "." + str[decimalPointIndex:]
	}

	str = trimInsignificantDigits(str)

	if s.neg {
		str = "-" + str
	}

	return str
}

func trimInsignificantDigits(str string) string {
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

	if str[0] == '.' {
		str = "0" + str
	}

	return str
}
