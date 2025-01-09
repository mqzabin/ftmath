package fuzzdecimal

import (
	"testing"

	"github.com/mqzabin/fuzzdecimal/fdlib"
)

// Fuzz initializes the fuzzing with the given numbersCount, and should be called once for each Fuzz...(f *testing.F) function.
//
// The numbersCount defines how many decimal seeds will be generated for the fuzzy process.
// Given a numbersCount of N, you can call the following package functions inside the fuzzFunc:
//
//   - AsStringN
//   - AsStringSlice
//   - AsDecimalN
//   - AsDecimalSlice
//   - AsDecimalComparisonN
//   - AsDecimalComparisonSlice
//
// Some options could be provided to customize the generated value individually (or all of them). The options are:
//
//   - Defines DecimalOption for a specific decimal index: WithDecimal(index int, options ...DecimalOption) Option
//   - Defines DecimalOption to be applied to all decimal indexes: WithAllDecimals(options ...DecimalOption) Option
//
// Options will be applied in order, so you can define a default value with WithAllDecimals, then specify some
// configuration for a specific decimal index.
func Fuzz(f *testing.F, numbersCount int, fuzzFunc func(t *T), options ...Option) {
	f.Helper()

	cfg := parseConfig(f, numbersCount, options)

	rvFunc := fdlib.CreateSeedFunc(f, cfg, fdlib.SeedsFuncToStringsFunc(f, cfg, func(t *testing.T, strNumbers []string) {
		t.Helper()

		fuzzFunc(&T{
			T:     t,
			seeds: strNumbers,
		})
	}))

	f.Fuzz(rvFunc.Interface())
}
