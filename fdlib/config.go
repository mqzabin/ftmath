package fdlib

import "testing"

const (
	DefaultDecimalMaxSignificantDigits = 10
	DefaultDecimalSigned               = true
	DefaultDecimalMaxDecimalPlaces     = 2
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
	MaxDecimalPlaces     int
}

func NewDecimalConfig(f *testing.F) DecimalConfig {
	f.Helper()

	return DecimalConfig{
		MaxSignificantDigits: DefaultDecimalMaxSignificantDigits,
		Signed:               DefaultDecimalSigned,
		MaxDecimalPlaces:     DefaultDecimalMaxDecimalPlaces,
	}
}

// Validate validates the configuration after all options have been applied.
// This serves to make the option order to be irrelevant.
func (cfg DecimalConfig) Validate(f *testing.F) {
	f.Helper()

	if cfg.MaxSignificantDigits < cfg.MaxDecimalPlaces {
		f.Fatalf("max decimal places %d cannot be greater than max significant digits %d", cfg.MaxDecimalPlaces, cfg.MaxSignificantDigits)
	}
}
