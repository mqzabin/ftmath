package fuzzdecimal

import (
	"testing"
)

func Fuzz(f *testing.F, numbersCount int, fuzzFunc func(t *T), options ...Option) {
	f.Helper()

	cfg := parseFuzzerConfig(f, numbersCount, options)

	rvFunc := createSeedFunc(f, cfg, seedToStringFunc(f, cfg, func(t *testing.T, strNumbers []string) {
		t.Helper()

		fuzzFunc(&T{
			T:     t,
			seeds: strNumbers,
		})
	}))

	f.Fuzz(rvFunc.Interface())
}
