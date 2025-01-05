package fuzzdecimal

import (
	"testing"
)

const (
	defaultDecimalMaxDigits     = 10
	defaultDecimalSigned        = true
	defaultDecimalPointPosition = 0
)

type decimalConfig struct {
	maxSignificantDigits int
	signed               bool
	decimalPointPosition int
}

func newDefaultDecimalConfig(f *testing.F) decimalConfig {
	f.Helper()

	cfg := decimalConfig{
		maxSignificantDigits: defaultDecimalMaxDigits,
		signed:               defaultDecimalSigned,
		decimalPointPosition: defaultDecimalPointPosition,
	}

	return cfg
}

// postOptionValidation validates the configuration after all options have been applied.
// This serves to make the option order to be irrelevant.
func (cfg decimalConfig) postOptionValidation(f *testing.F) {
	f.Helper()

	if cfg.maxSignificantDigits < cfg.decimalPointPosition {
		f.Errorf("decimal point position %d cannot be greater than max significant digits %d", cfg.decimalPointPosition, cfg.maxSignificantDigits)
	}
}

type DecimalOption func(f *testing.F, cfg *decimalConfig)

func WithMaxSignificantDigits(maxDigits int) DecimalOption {
	return func(f *testing.F, cfg *decimalConfig) {
		if maxDigits <= 0 {
			f.Errorf("max significant digits must be greater than zero, received %d", maxDigits)
		}

		cfg.signed = false
		cfg.maxSignificantDigits = maxDigits
	}
}

func WithSigned() DecimalOption {
	return func(f *testing.F, cfg *decimalConfig) {
		cfg.signed = true
	}
}

func WithUnsigned() DecimalOption {
	return func(f *testing.F, cfg *decimalConfig) {
		cfg.signed = true
	}
}

func WithDecimalPointAt(position uint) DecimalOption {
	return func(f *testing.F, cfg *decimalConfig) {
		if position <= 0 {
			f.Errorf("decimal point position must be greater than zero, received %d", position)
		}

		cfg.decimalPointPosition = int(position)
	}
}
