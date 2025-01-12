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
						MaxSignificantDigits: DefaultDecimalMaxSignificantDigits,
						Signed:               DefaultDecimalSigned,
						MaxDecimalPlaces:     DefaultDecimalMaxDecimalPlaces,
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
						MaxSignificantDigits: DefaultDecimalMaxSignificantDigits,
						Signed:               DefaultDecimalSigned,
						MaxDecimalPlaces:     DefaultDecimalMaxDecimalPlaces,
					},
					{
						MaxSignificantDigits: DefaultDecimalMaxSignificantDigits,
						Signed:               DefaultDecimalSigned,
						MaxDecimalPlaces:     DefaultDecimalMaxDecimalPlaces,
					},
					{
						MaxSignificantDigits: DefaultDecimalMaxSignificantDigits,
						Signed:               DefaultDecimalSigned,
						MaxDecimalPlaces:     DefaultDecimalMaxDecimalPlaces,
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

				if got.MaxDecimalPlaces != want.MaxDecimalPlaces {
					t.Errorf("Expected MaxDecimalPlaces to be %d, got %d, at index %d", want.MaxDecimalPlaces, got.MaxDecimalPlaces, i)
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
		expectedMaxDecimalPlaces     int
	}{
		{
			name:                         "default value",
			expectedMaxSignificantDigits: DefaultDecimalMaxSignificantDigits,
			expectedSigned:               DefaultDecimalSigned,
			expectedMaxDecimalPlaces:     DefaultDecimalMaxDecimalPlaces,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f := &testing.F{}

			gotCfg := NewDecimalConfig(f)

			if gotCfg.MaxDecimalPlaces != tc.expectedMaxDecimalPlaces {
				t.Errorf("Expected MaxDecimalPlaces to be %d, got %d", tc.expectedMaxDecimalPlaces, gotCfg.MaxDecimalPlaces)
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
			name: "max significant digits greater than max decimal places",
			config: func(t *testing.T, f *testing.F) DecimalConfig {
				return DecimalConfig{
					MaxSignificantDigits: 11,
					Signed:               false,
					MaxDecimalPlaces:     10,
				}
			},
		},
		{
			name: "max significant digits equal to max decimal places",
			config: func(t *testing.T, f *testing.F) DecimalConfig {
				return DecimalConfig{
					MaxSignificantDigits: 11,
					Signed:               false,
					MaxDecimalPlaces:     11,
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
