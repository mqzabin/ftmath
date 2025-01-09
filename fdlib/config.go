package fdlib

import "testing"

const (
	DefaultDecimalMaxDigits     = 10
	DefaultDecimalSigned        = true
	DefaultDecimalPointPosition = 2
)

type Config struct {
	Decimals []DecimalConfig
}

func NewConfig(f *testing.F, decimalsCount int) Config {
	f.Helper()

	if decimalsCount <= 0 {
		f.Fatalf("decimals count must be greater than zero, received %d", decimalsCount)
	}

	cfg := Config{
		Decimals: make([]DecimalConfig, 0, decimalsCount),
	}

	for range decimalsCount {
		cfg.Decimals = append(cfg.Decimals, NewDecimalConfig(f))
	}

	return cfg
}

type DecimalConfig struct {
	MaxSignificantDigits int
	Signed               bool
	DecimalPointPosition int
}

func NewDecimalConfig(f *testing.F) DecimalConfig {
	f.Helper()

	return DecimalConfig{
		MaxSignificantDigits: DefaultDecimalMaxDigits,
		Signed:               DefaultDecimalSigned,
		DecimalPointPosition: DefaultDecimalPointPosition,
	}
}

// Validate validates the configuration after all options have been applied.
// This serves to make the option order to be irrelevant.
func (cfg DecimalConfig) Validate(f *testing.F) {
	f.Helper()

	if cfg.MaxSignificantDigits < cfg.DecimalPointPosition {
		f.Fatalf("decimal point position %d cannot be greater than max significant digits %d", cfg.DecimalPointPosition, cfg.MaxSignificantDigits)
	}
}
