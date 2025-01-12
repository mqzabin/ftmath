package fdlib

import "testing"

func TestNewConfig(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		decimalsCount int
		expected      Config
	}{
		{
			name:          "1 decimals count",
			decimalsCount: 1,
			expected: Config{
				Decimals: []DecimalConfig{
					{
						MaxSignificantDigits: DefaultDecimalMaxDigits,
						Signed:               DefaultDecimalSigned,
						DecimalPointPosition: DefaultDecimalPointPosition,
					},
				},
			},
		},
		{
			name:          "3 decimals count",
			decimalsCount: 1,
			expected: Config{
				Decimals: []DecimalConfig{
					{
						MaxSignificantDigits: DefaultDecimalMaxDigits,
						Signed:               DefaultDecimalSigned,
						DecimalPointPosition: DefaultDecimalPointPosition,
					},
					{
						MaxSignificantDigits: DefaultDecimalMaxDigits,
						Signed:               DefaultDecimalSigned,
						DecimalPointPosition: DefaultDecimalPointPosition,
					},
					{
						MaxSignificantDigits: DefaultDecimalMaxDigits,
						Signed:               DefaultDecimalSigned,
						DecimalPointPosition: DefaultDecimalPointPosition,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f := &testing.F{}

			gotCfg := NewConfig(f, tc.decimalsCount)

			if len(gotCfg.Decimals) != tc.decimalsCount {
				t.Fatalf("Expected Decimals to have %d elements, got %d", tc.decimalsCount, len(gotCfg.Decimals))
			}

			for i := range gotCfg.Decimals {
				want, got := tc.expected.Decimals[i], gotCfg.Decimals[i]

				if got.DecimalPointPosition != want.DecimalPointPosition {
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

func TestNewDecimalConfig(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                         string
		expectedMaxSignificantDigits int
		expectedSigned               bool
		expectedDecimalPointPosition int
	}{
		{
			name:                         "default value",
			expectedMaxSignificantDigits: DefaultDecimalMaxDigits,
			expectedSigned:               DefaultDecimalSigned,
			expectedDecimalPointPosition: DefaultDecimalPointPosition,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f := &testing.F{}

			gotCfg := NewDecimalConfig(f)

			if gotCfg.DecimalPointPosition != tc.expectedDecimalPointPosition {
				t.Errorf("Expected DecimalPointPosition to be %d, got %d", tc.expectedDecimalPointPosition, gotCfg.DecimalPointPosition)
			}

			if gotCfg.MaxSignificantDigits != tc.expectedMaxSignificantDigits {
				t.Errorf("Expected MaxSignificantDigits to be %d, got %d", tc.expectedMaxSignificantDigits, gotCfg.MaxSignificantDigits)
			}

			if gotCfg.Signed != tc.expectedSigned {
				t.Errorf("Expected Signed to be %v, got %v", tc.expectedSigned, gotCfg.Signed)
			}
		})
	}
}

func TestDecimalConfigValidate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		config func(t *testing.T, f *testing.F) DecimalConfig
	}{
		{
			name: "default value",
			config: func(t *testing.T, f *testing.F) DecimalConfig {
				return NewDecimalConfig(f)
			},
		},
		{
			name: "max significant digits greater than decimal point position",
			config: func(t *testing.T, f *testing.F) DecimalConfig {
				return DecimalConfig{
					MaxSignificantDigits: 11,
					Signed:               false,
					DecimalPointPosition: 10,
				}
			},
		},
		{
			name: "max significant digits equal to decimal point position",
			config: func(t *testing.T, f *testing.F) DecimalConfig {
				return DecimalConfig{
					MaxSignificantDigits: 11,
					Signed:               false,
					DecimalPointPosition: 11,
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f := &testing.F{}

			tc.config(t, f).Validate(f)
		})
	}
}
