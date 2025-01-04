package fuzzdecimal

import (
	"testing"
)

const (
	defaultMaxDigits = 10
	defaultSigned    = true
)

type config struct {
	maxDigits      int
	uintsPerNumber int
	signed         bool
}

func parseConfig(f *testing.F, options []Option) config {
	f.Helper()

	cfg := config{
		maxDigits:      defaultMaxDigits,
		uintsPerNumber: UintsPerNumber(defaultMaxDigits),
		signed:         defaultSigned,
	}

	for _, opt := range options {
		opt(f, &cfg)
	}

	return cfg
}

type Option func(f *testing.F, cfg *config)

func WithUnsignedMaxDigits(maxDigits int) Option {
	return func(f *testing.F, cfg *config) {
		f.Helper()

		cfg.signed = false
		cfg.maxDigits = maxDigits
		cfg.uintsPerNumber = UintsPerNumber(maxDigits)
	}
}

func WithSignedMaxDigits(maxDigits int) Option {
	return func(f *testing.F, cfg *config) {
		f.Helper()

		cfg.signed = true
		cfg.maxDigits = maxDigits
		cfg.uintsPerNumber = UintsPerNumber(maxDigits)
	}
}
