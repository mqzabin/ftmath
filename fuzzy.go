package ftmath

import (
	"testing"
)

type Fuzzer[N any] struct {
	f         *testing.F
	parseFunc func(t *testing.T, s string) N
	cfg       config
}

func NewFuzzer[N any](f *testing.F, parseNumberFunc func(t *testing.T, s string) N, options ...Option) *Fuzzer[N] {
	f.Helper()

	cfg := parseConfig(f, options)

	return &Fuzzer[N]{
		f:         f,
		parseFunc: parseNumberFunc,
		cfg:       cfg,
	}
}

func (ft *Fuzzer[N]) fuzzN(numbersCount int, fuzzFunc func(t *testing.T, numbers []N)) {
	ft.f.Helper()

	rtFuncSignature := createFuncSignature(ft.f, ft.cfg.signed, numbersCount, ft.cfg.uintsPerNumber)

	seedAdapter := seedToStringAdapter(ft.f, ft.cfg.maxDigits, ft.parseFunc, fuzzFunc)

	rvFunc := createSeedFunc(ft.f, ft.cfg.signed, numbersCount, ft.cfg.uintsPerNumber, rtFuncSignature, seedAdapter)

	ft.f.Fuzz(rvFunc.Interface())
}

func (ft *Fuzzer[N]) Fuzz1(testFunc func(t *testing.T, x1 N)) {
	ft.f.Helper()

	ft.fuzzN(1, func(t *testing.T, numbers []N) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0])
	})
}

func (ft *Fuzzer[N]) Fuzz2(testFunc func(t *testing.T, x1, x2 N)) {
	ft.f.Helper()

	ft.fuzzN(2, func(t *testing.T, numbers []N) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1])
	})
}

func (ft *Fuzzer[N]) Fuzz3(testFunc func(t *testing.T, x1, x2, x3 N)) {
	ft.f.Helper()

	ft.fuzzN(3, func(t *testing.T, numbers []N) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2])
	})
}

func (ft *Fuzzer[N]) Fuzz4(testFunc func(t *testing.T, x1, x2, x3, x4 N)) {
	ft.f.Helper()

	ft.fuzzN(4, func(t *testing.T, numbers []N) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3])
	})
}

func (ft *Fuzzer[N]) Fuzz5(testFunc func(t *testing.T, x1, x2, x3, x4, x5 N)) {
	ft.f.Helper()

	ft.fuzzN(5, func(t *testing.T, numbers []N) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
	})
}

func (ft *Fuzzer[N]) Fuzz6(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6 N)) {
	ft.f.Helper()

	ft.fuzzN(6, func(t *testing.T, numbers []N) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
	})
}

func (ft *Fuzzer[N]) Fuzz7(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7 N)) {
	ft.f.Helper()

	ft.fuzzN(7, func(t *testing.T, numbers []N) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
	})
}

func (ft *Fuzzer[N]) Fuzz8(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8 N)) {
	ft.f.Helper()

	ft.fuzzN(8, func(t *testing.T, numbers []N) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
	})
}

func (ft *Fuzzer[N]) Fuzz9(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9 N)) {
	ft.f.Helper()

	ft.fuzzN(9, func(t *testing.T, numbers []N) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
	})
}
func (ft *Fuzzer[N]) Fuzz10(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 N)) {
	ft.f.Helper()

	ft.fuzzN(10, func(t *testing.T, numbers []N) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
	})
}
