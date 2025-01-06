package fuzzdecimal

import (
	"testing"
)

type config struct {
	decimals []decimalConfig
}

func parseFuzzerConfig(f *testing.F, decimalsCount int, options []Option) config {
	f.Helper()

	cfg := config{
		decimals: make([]decimalConfig, 0, decimalsCount),
	}

	// Initializing number configs.
	for range decimalsCount {
		cfg.decimals = append(cfg.decimals, newDefaultDecimalConfig(f))
	}

	for _, opt := range options {
		opt(f, &cfg)
	}

	for _, decCfg := range cfg.decimals {
		decCfg.postOptionValidation(f)
	}

	return cfg
}

func (cfg config) applyDecimalOption(f *testing.F, numberIndex int, options []DecimalOption) {
	f.Helper()

	if numberIndex <= 0 {
		f.Errorf("decimals are 1-indexed, received %d index value", numberIndex)
	}

	sliceIndex := numberIndex - 1

	if sliceIndex > len(cfg.decimals) {
		f.Errorf("decimals count is %d, received %d index value", len(cfg.decimals), numberIndex)
	}

	for _, opt := range options {
		opt(f, &cfg.decimals[sliceIndex])
	}
}

// Option represents an option for the Fuzz function call. The available options are:
//
//   - WithDecimal(index int, options ...DecimalOption) Option
//   - WithAllDecimals(options ...DecimalOption) Option
type Option func(f *testing.F, cfg *config)

// WithAllDecimals defines DecimalOptions for all generated decimal seed values.
// Check DecimalOption for available sub-options.
func WithAllDecimals(options ...DecimalOption) Option {
	return func(f *testing.F, cfg *config) {
		f.Helper()

		for i := range len(cfg.decimals) {
			cfg.applyDecimalOption(f, i+1, options)
		}
	}
}

// WithDecimal defines DecimalOptions for a specific decimal seed value at the provided index.
// The index should be 1-indexes, i.e. the first seed is at index 1.
// Check DecimalOption for available sub-options.
func WithDecimal(index int, options ...DecimalOption) Option {
	return func(f *testing.F, cfg *config) {
		f.Helper()

		cfg.applyDecimalOption(f, index, options)
	}
}
