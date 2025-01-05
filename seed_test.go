package fuzzdecimal

import (
	"testing"
)

func Test_seed_string(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		seed   seed
		config decimalConfig
		want   string
	}{
		{
			name: "no uints",
			seed: seed{
				uints: nil,
				neg:   false,
			},
			config: decimalConfig{
				maxSignificantDigits: 30,
				signed:               true,
				decimalPointPosition: 15,
			},
			want: "0",
		},
		{
			name: "single zero uint",
			seed: seed{
				uints: []uint64{0},
				neg:   false,
			},
			config: decimalConfig{
				maxSignificantDigits: 30,
				signed:               true,
				decimalPointPosition: 15,
			},
			want: "0",
		},
		{
			name: "many uints, all zeroes",
			seed: seed{
				uints: []uint64{0, 0, 0, 0, 0},
				neg:   false,
			},
			config: decimalConfig{
				maxSignificantDigits: 30,
				signed:               true,
				decimalPointPosition: 15,
			},
			want: "0",
		},
		{
			name: "omit zero negative sign",
			seed: seed{
				uints: []uint64{0, 0, 0, 0, 0},
				neg:   true,
			},
			config: decimalConfig{
				maxSignificantDigits: 30,
				signed:               true,
				decimalPointPosition: 15,
			},
			want: "0",
		},
		{
			name: "single maxUintValue uint",
			seed: seed{
				uints: []uint64{maxUintValue},
				neg:   false,
			},
			config: decimalConfig{
				maxSignificantDigits: 30,
				signed:               true,
				decimalPointPosition: 0,
			},
			want: "9999999999999999999",
		},
		{
			name: "multiple maxUintValue uint with decimal pointer",
			seed: seed{
				uints: []uint64{maxUintValue, maxUintValue},
				neg:   false,
			},
			config: decimalConfig{
				maxSignificantDigits: 38,
				signed:               true,
				decimalPointPosition: 15,
			},
			want: "99999999999999999999999.999999999999999",
		},
		{
			name: "multiple maxUintValue uint with decimal point on last index",
			seed: seed{
				uints: []uint64{maxUintValue, maxUintValue},
				neg:   false,
			},
			config: decimalConfig{
				maxSignificantDigits: 38,
				signed:               true,
				decimalPointPosition: 0,
			},
			want: "99999999999999999999999999999999999999",
		},
		{
			name: "multiple maxUintValue uint with decimal point on first index",
			seed: seed{
				uints: []uint64{maxUintValue, maxUintValue},
				neg:   false,
			},
			config: decimalConfig{
				maxSignificantDigits: 38,
				signed:               true,
				decimalPointPosition: 38,
			},
			want: "0.99999999999999999999999999999999999999",
		},
		{
			name: "minimum non-zero",
			seed: seed{
				uints: []uint64{0, 1},
				neg:   false,
			},
			config: decimalConfig{
				maxSignificantDigits: 38,
				signed:               true,
				decimalPointPosition: 15,
			},
			want: "0.000000000000001",
		},
		{
			name: "minimum non-zero with max decimal point position",
			seed: seed{
				uints: []uint64{0, 1},
				neg:   false,
			},
			config: decimalConfig{
				maxSignificantDigits: 38,
				signed:               true,
				decimalPointPosition: 38,
			},
			want: "0.00000000000000000000000000000000000001",
		},
		{
			name: "negative minimum non-zero with max decimal point position",
			seed: seed{
				uints: []uint64{0, 1},
				neg:   true,
			},
			config: decimalConfig{
				maxSignificantDigits: 38,
				signed:               true,
				decimalPointPosition: 38,
			},
			want: "-0.00000000000000000000000000000000000001",
		},
		{
			name: "negative minimum non-zero with max decimal point position",
			seed: seed{
				uints: []uint64{0, 0, 0, 48},
				neg:   false,
			},
			config: decimalConfig{
				maxSignificantDigits: 72,
				signed:               true,
				decimalPointPosition: 18,
			},
			want: "0.000000000000000048",
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.seed.string(tc.config)

			if got != tc.want {
				t.Errorf("seed.string() = '%s', want '%s'", got, tc.want)
			}
		})
	}
}
