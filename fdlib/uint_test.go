package fdlib

import (
	"testing"
)

func TestUintsPerNumber(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		maxDigits int
		want      int
	}{
		{
			name:      "0 digits",
			maxDigits: 0,
			want:      0,
		},
		{
			name:      "10 digits",
			maxDigits: 10,
			want:      1,
		},
		{
			name:      "19 digits",
			maxDigits: MaxDigitsPerUint,
			want:      1,
		},
		{
			name:      "20 digits",
			maxDigits: MaxDigitsPerUint + 1,
			want:      2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := UintsPerNumber(tc.maxDigits); got != tc.want {
				t.Errorf("UintsPerNumber() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestNormalizeUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		n         uint64
		maxDigits int
		want      uint64
	}{
		{
			name:      "overflows MaxDigitsPerUint digits",
			n:         18446744073709551615, // math.MathUint64
			maxDigits: MaxDigitsPerUint,
			want:      8446744073709551615,
		},
		{
			name:      "lower than MaxDigitsPerUint digits",
			n:         73709551615, // math.MathUint64
			maxDigits: MaxDigitsPerUint,
			want:      73709551615, // unchanged
		},
		{
			name:      "overflows 5 digits",
			n:         987654321,
			maxDigits: 5,
			want:      54321,
		},
		{
			name:      "lower than 5 digits",
			n:         321,
			maxDigits: 5,
			want:      321, // unchanged
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NormalizeUint(t, tc.n, tc.maxDigits); got != tc.want {
				t.Errorf("NormalizeUint() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestUintToString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		n    uint64
		want string
	}{
		{
			name: "0",
			n:    0,
			want: "0000000000000000000",
		},
		{
			name: "99",
			n:    99,
			want: "0000000000000000099",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := UintToString(tc.n); got != tc.want {
				t.Errorf("UintToString() = %s, want %s", got, tc.want)
			}
		})
	}
}
