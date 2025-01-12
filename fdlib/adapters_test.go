package fdlib

import (
	"testing"

	shopspring "github.com/shopspring/decimal"
)

func TestSeedsFuncToStringsFunc(t *testing.T) {
	t.Parallel()

	assertStrs := func(t *testing.T, expected, got []string) {
		t.Helper()

		if len(got) != len(expected) {
			t.Fatalf("expected %d strings, got %d", len(expected), len(got))
		}

		for i := range got {
			if got[i] != expected[i] {
				t.Fatalf("expected string '%s', got '%s'", expected[i], got[i])
			}
		}
	}
	assertStrs(t, nil, nil)

	testCases := []struct {
		name    string
		cfgFunc func(t *testing.T, f *testing.F) Config
		seeds   []Seed
		// Assert inside handler.
		handler func(t *testing.T, strNumbers []string)
	}{
		{
			name: "single zero seed",
			cfgFunc: func(t *testing.T, f *testing.F) Config {
				return NewConfig(f, 1)
			},
			seeds: []Seed{
				{
					Uints: nil,
					Neg:   false,
				},
			},
			handler: func(t *testing.T, strNumbers []string) {
				assertStrs(t, []string{"0"}, strNumbers)
			},
		},
		{
			name: "multiple seeds",
			cfgFunc: func(t *testing.T, f *testing.F) Config {
				decCfg := DecimalConfig{
					MaxSignificantDigits: MaxDigitsPerUint * 4,
					Signed:               true,
					DecimalPointPosition: 10,
				}

				return Config{
					Decimals: []DecimalConfig{decCfg, decCfg, decCfg},
				}
			},
			seeds: []Seed{
				{
					Uints: []uint64{0, 0, 1, 0},
					Neg:   false,
				},
				{
					Uints: []uint64{0, 1, 0, 0},
					Neg:   false,
				},
				{
					Uints: []uint64{0, 0, 0, 0},
					Neg:   false,
				},
			},
			handler: func(t *testing.T, strNumbers []string) {
				assertStrs(t, []string{"1000000000", "10000000000000000000000000000", "0"}, strNumbers)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f := &testing.F{}

			newFunc := SeedsFuncToStringsFunc(f, tc.cfgFunc(t, f), tc.handler)

			newFunc(t, tc.seeds)
		})
	}
}

func TestParseStringSliceToDecimalSlice(t *testing.T) {
	t.Parallel()

	mustDecimals := func(t *testing.T, strs ...string) []shopspring.Decimal {
		t.Helper()

		decs := make([]shopspring.Decimal, 0, len(strs))

		for _, s := range strs {
			dec, err := shopspring.NewFromString(s)
			if err != nil {
				t.Fatalf("failed to create decimal from string '%s': %v", s, err)
			}

			decs = append(decs, dec)
		}

		return decs
	}

	testCases := []struct {
		name             string
		decimalStrings   []string
		parseDecimalFunc func(t *testing.T, s string) (shopspring.Decimal, error)
		expectedDecimals []shopspring.Decimal
	}{
		{
			name:           "0 decimals",
			decimalStrings: []string{},
			parseDecimalFunc: func(t *testing.T, s string) (shopspring.Decimal, error) {
				return shopspring.NewFromString(s)
			},
			expectedDecimals: []shopspring.Decimal{},
		},
		{
			name: "1 zero decimals",
			decimalStrings: []string{
				"0",
			},
			parseDecimalFunc: func(t *testing.T, s string) (shopspring.Decimal, error) {
				return shopspring.NewFromString(s)
			},
			expectedDecimals: mustDecimals(t,
				"0",
			),
		},
		{
			name: "3 decimals",
			decimalStrings: []string{
				"1.23",
				"0",
				"91239123912992312",
			},
			parseDecimalFunc: func(t *testing.T, s string) (shopspring.Decimal, error) {
				return shopspring.NewFromString(s)
			},
			expectedDecimals: mustDecimals(t,
				"1.23",
				"0",
				"91239123912992312",
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotDecimals, err := ParseStringSliceToDecimalSlice(t, tc.decimalStrings, tc.parseDecimalFunc)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(gotDecimals) != len(tc.expectedDecimals) {
				t.Fatalf("expected %d decimals, got %d", len(tc.expectedDecimals), len(gotDecimals))
			}

			for i := range gotDecimals {
				if !gotDecimals[i].Equal(tc.expectedDecimals[i]) {
					t.Fatalf("expected decimal %s, got %s", tc.expectedDecimals[i].String(), gotDecimals[i].String())
				}
			}
		})
	}
}
