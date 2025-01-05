package fuzzdecimal

import (
	"testing"
)

func Test_uintsPerNumber(t *testing.T) {
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
			maxDigits: maxDigitsPerUint,
			want:      1,
		},
		{
			name:      "20 digits",
			maxDigits: maxDigitsPerUint + 1,
			want:      2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := uintsPerNumber(tc.maxDigits); got != tc.want {
				t.Errorf("uintsPerNumber() = %d, want %d", got, tc.want)
			}
		})
	}
}

func Test_normalizeUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		n         uint64
		maxDigits int
		want      uint64
	}{
		{
			name:      "overflows maxDigitsPerUint digits",
			n:         18446744073709551615, // math.MathUint64
			maxDigits: maxDigitsPerUint,
			want:      8446744073709551615,
		},
		{
			name:      "lower than maxDigitsPerUint digits",
			n:         73709551615, // math.MathUint64
			maxDigits: maxDigitsPerUint,
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
			if got := normalizeUint(tc.n, tc.maxDigits); got != tc.want {
				t.Errorf("normalizeUint() = %d, want %d", got, tc.want)
			}
		})
	}
}

func Test_uintToString(t *testing.T) {
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
			if got := uintToString(tc.n); got != tc.want {
				t.Errorf("uintToString() = %s, want %s", got, tc.want)
			}
		})
	}
}
