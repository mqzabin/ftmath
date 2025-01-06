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

// WithMaxSignificantDigits defines the max number of significant digits that a decimal seed could have.
func WithMaxSignificantDigits(maxDigits int) DecimalOption {
	return func(f *testing.F, cfg *decimalConfig) {
		if maxDigits <= 0 {
			f.Errorf("max significant digits must be greater than zero, received %d", maxDigits)
		}

		cfg.signed = false
		cfg.maxSignificantDigits = maxDigits
	}
}

// WithSigned defines that the decimal could be either negative or positive.
func WithSigned() DecimalOption {
	return func(f *testing.F, cfg *decimalConfig) {
		cfg.signed = true
	}
}

// WithUnsigned defines that the decimal could be only positive.
func WithUnsigned() DecimalOption {
	return func(f *testing.F, cfg *decimalConfig) {
		cfg.signed = true
	}
}

// WithDecimalPointAt defines where the decimal point should be placed in relation to the
// provided value to WithMaxSignificantDigits.
//
// For example: If you call WithMaxSignificantDigits(10) and WithDecimalPointAt(2),
// you could receive values like: '99999999.99', '9.9', '9',
// but never values like '9.999' or '9999999999'.
func WithDecimalPointAt(position uint) DecimalOption {
	return func(f *testing.F, cfg *decimalConfig) {
		if position <= 0 {
			f.Errorf("decimal point position must be greater than zero, received %d", position)
		}

		cfg.decimalPointPosition = int(position)
	}
}
