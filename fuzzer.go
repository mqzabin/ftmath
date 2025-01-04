package fuzzmath

import (
	"testing"
)

type Fuzzer struct {
	f   *testing.F
	cfg config
}

func NewFuzzer(f *testing.F, options ...Option) *Fuzzer {
	f.Helper()

	cfg := parseConfig(f, options)

	return &Fuzzer{
		f:   f,
		cfg: cfg,
	}
}

func (ft *Fuzzer) fuzzN(numbersCount int, fuzzFunc func(t *testing.T, numbers []string)) {
	ft.f.Helper()

	rtFuncSignature := createFuncSignature(ft.f, ft.cfg.signed, numbersCount, ft.cfg.uintsPerNumber)

	seedAdapter := seedToStringAdapter(ft.f, ft.cfg.maxDigits, fuzzFunc)

	rvFunc := createSeedFunc(ft.f, ft.cfg.signed, numbersCount, ft.cfg.uintsPerNumber, rtFuncSignature, seedAdapter)

	ft.f.Fuzz(rvFunc.Interface())
}

func (ft *Fuzzer) Fuzz1(testFunc func(t *testing.T, x1 string)) {
	ft.f.Helper()

	ft.fuzzN(1, func(t *testing.T, numbers []string) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0])
	})
}

func (ft *Fuzzer) Fuzz2(testFunc func(t *testing.T, x1, x2 string)) {
	ft.f.Helper()

	ft.fuzzN(2, func(t *testing.T, numbers []string) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1])
	})
}

func (ft *Fuzzer) Fuzz3(testFunc func(t *testing.T, x1, x2, x3 string)) {
	ft.f.Helper()

	ft.fuzzN(3, func(t *testing.T, numbers []string) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2])
	})
}

func (ft *Fuzzer) Fuzz4(testFunc func(t *testing.T, x1, x2, x3, x4 string)) {
	ft.f.Helper()

	ft.fuzzN(4, func(t *testing.T, numbers []string) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3])
	})
}

func (ft *Fuzzer) Fuzz5(testFunc func(t *testing.T, x1, x2, x3, x4, x5 string)) {
	ft.f.Helper()

	ft.fuzzN(5, func(t *testing.T, numbers []string) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
	})
}

func (ft *Fuzzer) Fuzz6(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6 string)) {
	ft.f.Helper()

	ft.fuzzN(6, func(t *testing.T, numbers []string) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
	})
}

func (ft *Fuzzer) Fuzz7(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7 string)) {
	ft.f.Helper()

	ft.fuzzN(7, func(t *testing.T, numbers []string) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
	})
}

func (ft *Fuzzer) Fuzz8(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8 string)) {
	ft.f.Helper()

	ft.fuzzN(8, func(t *testing.T, numbers []string) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
	})
}

func (ft *Fuzzer) Fuzz9(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9 string)) {
	ft.f.Helper()

	ft.fuzzN(9, func(t *testing.T, numbers []string) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
	})
}
func (ft *Fuzzer) Fuzz10(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 string)) {
	ft.f.Helper()

	ft.fuzzN(10, func(t *testing.T, numbers []string) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
	})
}
