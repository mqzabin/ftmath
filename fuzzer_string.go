package fuzzdecimal

import "testing"

type StringFuzzer struct {
	f   *testing.F
	cfg config
}

func NewStringFuzzer(f *testing.F, options ...Option) *StringFuzzer {
	f.Helper()

	cfg := parseConfig(f, options)

	return &StringFuzzer{
		f:   f,
		cfg: cfg,
	}
}

func (sf *StringFuzzer) fuzzN(numbersCount int, fuzzFunc func(t *testing.T, numbers []string)) {
	seedHandlerFunc := seedToStringAdapter(sf.cfg.maxDigits, fuzzFunc)

	rvFunc := createSeedFunc(sf.f, sf.cfg.signed, numbersCount, sf.cfg.uintsPerNumber, seedHandlerFunc)

	sf.f.Fuzz(rvFunc.Interface())
}

func (sf *StringFuzzer) Fuzz1(testFunc func(t *testing.T, x1 string)) {
	sf.fuzzN(1, func(t *testing.T, numbers []string) {
		t.Helper()

		testFunc(t, numbers[0])
	})
}

func (sf *StringFuzzer) Fuzz2(testFunc func(t *testing.T, x1, x2 string)) {
	sf.fuzzN(2, func(t *testing.T, numbers []string) {
		t.Helper()

		testFunc(t, numbers[0], numbers[1])
	})
}

func (sf *StringFuzzer) Fuzz3(testFunc func(t *testing.T, x1, x2, x3 string)) {
	sf.fuzzN(3, func(t *testing.T, numbers []string) {
		t.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2])
	})
}

func (sf *StringFuzzer) Fuzz4(testFunc func(t *testing.T, x1, x2, x3, x4 string)) {
	sf.fuzzN(4, func(t *testing.T, numbers []string) {
		t.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3])
	})
}

func (sf *StringFuzzer) Fuzz5(testFunc func(t *testing.T, x1, x2, x3, x4, x5 string)) {
	sf.fuzzN(5, func(t *testing.T, numbers []string) {
		t.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
	})
}

func (sf *StringFuzzer) Fuzz6(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6 string)) {
	sf.fuzzN(6, func(t *testing.T, numbers []string) {
		t.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
	})
}

func (sf *StringFuzzer) Fuzz7(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7 string)) {
	sf.fuzzN(7, func(t *testing.T, numbers []string) {
		t.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
	})
}

func (sf *StringFuzzer) Fuzz8(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8 string)) {
	sf.fuzzN(8, func(t *testing.T, numbers []string) {
		t.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
	})
}

func (sf *StringFuzzer) Fuzz9(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9 string)) {
	sf.fuzzN(9, func(t *testing.T, numbers []string) {
		t.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
	})
}

func (sf *StringFuzzer) Fuzz10(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 string)) {
	sf.fuzzN(10, func(t *testing.T, numbers []string) {
		t.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
	})
}
