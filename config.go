package fuzzmath

import (
	"testing"
)

const (
	defaultMaxDigits = uintSafeDigits
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
		uintsPerNumber: defaultMaxDigits / uintSafeDigits,
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
		cfg.uintsPerNumber = maxDigits / uintSafeDigits

	}
}

func WithSignedMaxDigits(maxDigits int) Option {
	return func(f *testing.F, cfg *config) {
		f.Helper()

		cfg.signed = true
		cfg.maxDigits = maxDigits
		cfg.uintsPerNumber = maxDigits / uintSafeDigits
	}
}
