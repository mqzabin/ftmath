package fuzzdecimal

import (
	"testing"

	"github.com/mqzabin/fuzzdecimal/fdlib"
)

func TestOptions(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		decimalsCount  int
		options        []Option
		expectedConfig fdlib.Config
	}{
		{
			name:          "no options",
			decimalsCount: 3,
			options:       []Option{},
			expectedConfig: fdlib.Config{
				Decimals: []fdlib.DecimalConfig{
					{
						MaxSignificantDigits: fdlib.DefaultDecimalMaxDigits,
						Signed:               fdlib.DefaultDecimalSigned,
						DecimalPointPosition: fdlib.DefaultDecimalPointPosition,
					},
					{
						MaxSignificantDigits: fdlib.DefaultDecimalMaxDigits,
						Signed:               fdlib.DefaultDecimalSigned,
						DecimalPointPosition: fdlib.DefaultDecimalPointPosition,
					},
					{
						MaxSignificantDigits: fdlib.DefaultDecimalMaxDigits,
						Signed:               fdlib.DefaultDecimalSigned,
						DecimalPointPosition: fdlib.DefaultDecimalPointPosition,
					},
				},
			},
		},
		{
			name:          "options for indexed decimals",
			decimalsCount: 3,
			options: []Option{
				WithDecimal(1, WithMaxSignificantDigits(20), WithDecimalPointAt(10), WithSigned()),
				WithDecimal(3, WithMaxSignificantDigits(5), WithDecimalPointAt(1), WithUnsigned()),
			},
			expectedConfig: fdlib.Config{
				Decimals: []fdlib.DecimalConfig{
					{
						MaxSignificantDigits: 20,
						Signed:               true,
						DecimalPointPosition: 10,
					},
					{
						MaxSignificantDigits: fdlib.DefaultDecimalMaxDigits,
						Signed:               fdlib.DefaultDecimalSigned,
						DecimalPointPosition: fdlib.DefaultDecimalPointPosition,
					},
					{
						MaxSignificantDigits: 5,
						Signed:               false,
						DecimalPointPosition: 1,
					},
				},
			},
		},
		{
			name:          "options for indexed decimals with indexed overrides",
			decimalsCount: 3,
			options: []Option{
				WithDecimal(1, WithMaxSignificantDigits(20), WithDecimalPointAt(10), WithSigned()),
				WithDecimal(3, WithMaxSignificantDigits(5), WithDecimalPointAt(1), WithUnsigned()),
				WithDecimal(3, WithSigned()),
			},
			expectedConfig: fdlib.Config{
				Decimals: []fdlib.DecimalConfig{
					{
						MaxSignificantDigits: 20,
						Signed:               true,
						DecimalPointPosition: 10,
					},
					{
						MaxSignificantDigits: fdlib.DefaultDecimalMaxDigits,
						Signed:               fdlib.DefaultDecimalSigned,
						DecimalPointPosition: fdlib.DefaultDecimalPointPosition,
					},
					{
						MaxSignificantDigits: 5,
						Signed:               true,
						DecimalPointPosition: 1,
					},
				},
			},
		},
		{
			name:          "global decimal options",
			decimalsCount: 3,
			options: []Option{
				WithAllDecimals(WithMaxSignificantDigits(20), WithDecimalPointAt(10), WithUnsigned()),
			},
			expectedConfig: fdlib.Config{
				Decimals: []fdlib.DecimalConfig{
					{
						MaxSignificantDigits: 20,
						Signed:               false,
						DecimalPointPosition: 10,
					},
					{
						MaxSignificantDigits: 20,
						Signed:               false,
						DecimalPointPosition: 10,
					},
					{
						MaxSignificantDigits: 20,
						Signed:               false,
						DecimalPointPosition: 10,
					},
				},
			},
		},
		{
			name:          "global decimals options with global overrides",
			decimalsCount: 3,
			options: []Option{
				WithAllDecimals(WithMaxSignificantDigits(20), WithDecimalPointAt(10), WithUnsigned()),
				WithAllDecimals(WithSigned()),
			},
			expectedConfig: fdlib.Config{
				Decimals: []fdlib.DecimalConfig{
					{
						MaxSignificantDigits: 20,
						Signed:               true,
						DecimalPointPosition: 10,
					},
					{
						MaxSignificantDigits: 20,
						Signed:               true,
						DecimalPointPosition: 10,
					},
					{
						MaxSignificantDigits: 20,
						Signed:               true,
						DecimalPointPosition: 10,
					},
				},
			},
		},
		{
			name:          "global options with indexed overrides",
			decimalsCount: 3,
			options: []Option{
				WithAllDecimals(WithMaxSignificantDigits(20), WithDecimalPointAt(10), WithUnsigned()),
				WithDecimal(2, WithSigned()),
			},
			expectedConfig: fdlib.Config{
				Decimals: []fdlib.DecimalConfig{
					{
						MaxSignificantDigits: 20,
						Signed:               false,
						DecimalPointPosition: 10,
					},
					{
						MaxSignificantDigits: 20,
						Signed:               true,
						DecimalPointPosition: 10,
					},
					{
						MaxSignificantDigits: 20,
						Signed:               false,
						DecimalPointPosition: 10,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f := &testing.F{}

			gotCfg := parseConfig(f, tc.decimalsCount, tc.options)

			if len(gotCfg.Decimals) != len(tc.expectedConfig.Decimals) {
				t.Fatalf("Expected Decimals to have %d elements, got %d", len(tc.expectedConfig.Decimals), len(gotCfg.Decimals))
			}

			for i := range gotCfg.Decimals {
				want, got := tc.expectedConfig.Decimals[i], gotCfg.Decimals[i]

				if gotCfg.Decimals[i].DecimalPointPosition != want.DecimalPointPosition {
					t.Errorf("Expected DecimalPointPosition to be %d, got %d, at index %d", want.DecimalPointPosition, got.DecimalPointPosition, i)
				}

				if got.MaxSignificantDigits != want.MaxSignificantDigits {
					t.Errorf("Expected MaxSignificantDigits to be %d, got %d, at index %d", want.MaxSignificantDigits, got.MaxSignificantDigits, i)
				}

				if got.Signed != want.Signed {
					t.Errorf("Expected Signed to be %v, got %v, at index %d", want.Signed, got.Signed, i)
				}
			}
		})
	}
}
