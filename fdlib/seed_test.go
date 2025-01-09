package fdlib

import "testing"

func TestSeedString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		seed   Seed
		config DecimalConfig
		want   string
	}{
		{
			name: "no uints",
			seed: Seed{
				Uints: nil,
				Neg:   false,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 30,
				Signed:               true,
				DecimalPointPosition: 15,
			},
			want: "0",
		},
		{
			name: "single zero uint",
			seed: Seed{
				Uints: []uint64{0},
				Neg:   false,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 30,
				Signed:               true,
				DecimalPointPosition: 15,
			},
			want: "0",
		},
		{
			name: "many uints, all zeroes",
			seed: Seed{
				Uints: []uint64{0, 0, 0, 0, 0},
				Neg:   false,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 30,
				Signed:               true,
				DecimalPointPosition: 15,
			},
			want: "0",
		},
		{
			name: "omit zero negative sign",
			seed: Seed{
				Uints: []uint64{0, 0, 0, 0, 0},
				Neg:   true,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 30,
				Signed:               true,
				DecimalPointPosition: 15,
			},
			want: "0",
		},
		{
			name: "single MaxUintValue uint",
			seed: Seed{
				Uints: []uint64{MaxUintValue},
				Neg:   false,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 30,
				Signed:               true,
				DecimalPointPosition: 0,
			},
			want: "9999999999999999999",
		},
		{
			name: "multiple MaxUintValue uint with decimal pointer",
			seed: Seed{
				Uints: []uint64{MaxUintValue, MaxUintValue},
				Neg:   false,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 38,
				Signed:               true,
				DecimalPointPosition: 15,
			},
			want: "99999999999999999999999.999999999999999",
		},
		{
			name: "multiple MaxUintValue uint with decimal point on last index",
			seed: Seed{
				Uints: []uint64{MaxUintValue, MaxUintValue},
				Neg:   false,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 38,
				Signed:               true,
				DecimalPointPosition: 0,
			},
			want: "99999999999999999999999999999999999999",
		},
		{
			name: "multiple MaxUintValue uint with decimal point on first index",
			seed: Seed{
				Uints: []uint64{MaxUintValue, MaxUintValue},
				Neg:   false,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 38,
				Signed:               true,
				DecimalPointPosition: 38,
			},
			want: "0.99999999999999999999999999999999999999",
		},
		{
			name: "minimum non-zero",
			seed: Seed{
				Uints: []uint64{0, 1},
				Neg:   false,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 38,
				Signed:               true,
				DecimalPointPosition: 15,
			},
			want: "0.000000000000001",
		},
		{
			name: "minimum non-zero with max decimal point position",
			seed: Seed{
				Uints: []uint64{0, 1},
				Neg:   false,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 38,
				Signed:               true,
				DecimalPointPosition: 38,
			},
			want: "0.00000000000000000000000000000000000001",
		},
		{
			name: "negative minimum non-zero with max decimal point position",
			seed: Seed{
				Uints: []uint64{0, 1},
				Neg:   true,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 38,
				Signed:               true,
				DecimalPointPosition: 38,
			},
			want: "-0.00000000000000000000000000000000000001",
		},
		{
			name: "negative minimum non-zero with max decimal point position",
			seed: Seed{
				Uints: []uint64{0, 0, 0, 48},
				Neg:   false,
			},
			config: DecimalConfig{
				MaxSignificantDigits: 72,
				Signed:               true,
				DecimalPointPosition: 18,
			},
			want: "0.000000000000000048",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.seed.String(tc.config)

			if got != tc.want {
				t.Errorf("Seed.String() = '%s', want '%s'", got, tc.want)
			}
		})
	}
}
