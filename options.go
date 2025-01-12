package fuzzdecimal

import (
	"testing"

	"github.com/mqzabin/fuzzdecimal/fdlib"
)

// parseConfig parses the provided Option slice and returns a lib.Config instance.
func parseConfig(f *testing.F, decimalsCount int, options []Option) fdlib.Config {
	f.Helper()

	cfg := fdlib.NewConfig(f, decimalsCount)

	for _, opt := range options {
		opt(f, &cfg)
	}

	for _, decimalCfg := range cfg.Decimals {
		decimalCfg.Validate(f)
	}

	return cfg
}

// applyDecimalOption applies the provided DecimalOption to the decimal seed value at the provided index.
func applyDecimalOption(f *testing.F, cfg *fdlib.Config, index int, options []DecimalOption) {
	f.Helper()

	if index <= 0 {
		f.Fatalf("decimals are 1-indexed, received %d index value", index)
	}

	sliceIndex := index - 1

	if sliceIndex > len(cfg.Decimals) {
		f.Fatalf("decimals count is %d, received %d index value", len(cfg.Decimals), index)
	}

	for _, opt := range options {
		opt(f, &cfg.Decimals[sliceIndex])
	}
}

// Option represents an option for the Fuzz function call. The available options are:
//
//   - WithDecimal(index int, options ...DecimalOption) Option
//   - WithAllDecimals(options ...DecimalOption) Option
type Option func(f *testing.F, cfg *fdlib.Config)

// WithAllDecimals defines DecimalOptions for all generated decimal seed values.
// Check DecimalOption for available sub-options.
func WithAllDecimals(options ...DecimalOption) Option {
	return func(f *testing.F, cfg *fdlib.Config) {
		f.Helper()

		for i := range len(cfg.Decimals) {
			applyDecimalOption(f, cfg, i+1, options)
		}
	}
}

// WithDecimal defines DecimalOptions for a specific decimal seed value at the provided index.
// The index should be 1-indexes, i.e. the first seed is at index 1.
// Check DecimalOption for available sub-options.
func WithDecimal(index int, options ...DecimalOption) Option {
	return func(f *testing.F, cfg *fdlib.Config) {
		f.Helper()

		applyDecimalOption(f, cfg, index, options)
	}
}

type DecimalOption func(f *testing.F, cfg *fdlib.DecimalConfig)

// WithMaxSignificantDigits defines the max number of significant digits that a decimal seed could have.
func WithMaxSignificantDigits(maxDigits int) DecimalOption {
	return func(f *testing.F, cfg *fdlib.DecimalConfig) {
		f.Helper()

		if maxDigits <= 0 {
			f.Fatalf("max significant digits must be greater than zero, received %d", maxDigits)
		}

		cfg.MaxSignificantDigits = maxDigits
	}
}

// WithSigned defines that the decimal could be either negative or positive.
func WithSigned() DecimalOption {
	return func(f *testing.F, cfg *fdlib.DecimalConfig) {
		f.Helper()

		cfg.Signed = true
	}
}

// WithUnsigned defines that the decimal could be only positive.
func WithUnsigned() DecimalOption {
	return func(f *testing.F, cfg *fdlib.DecimalConfig) {
		f.Helper()

		cfg.Signed = false
	}
}

// WithMaxDecimalPlaces defines the max after decimal point digits.
//
// For example: If you call WithMaxSignificantDigits(10) and WithMaxDecimalPlaces(2),
// you could receive values like: '99999999.99', '9.9', '9',
// but never values like '9.999' or '9999999999'.
func WithMaxDecimalPlaces(places uint) DecimalOption {
	return func(f *testing.F, cfg *fdlib.DecimalConfig) {
		f.Helper()

		if places < 0 {
			f.Fatalf("max decimal places must be greater than zero, received %d", places)
		}

		cfg.MaxDecimalPlaces = int(places)
	}
}
