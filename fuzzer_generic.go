package fuzzdecimal

import "testing"

type GenericFuzzer[Number any] struct {
	f               *testing.F
	parseNumberFunc func(t *testing.F, s string) (Number, error)
	cfg             config
}

func NewGenericFuzzer[Number any](f *testing.F, parseNumberFunc func(f *testing.F, s string) (Number, error), options ...Option) *GenericFuzzer[Number] {
	f.Helper()

	cfg := parseConfig(f, options)

	return &GenericFuzzer[Number]{
		f:               f,
		parseNumberFunc: parseNumberFunc,
		cfg:             cfg,
	}
}

func (ft *GenericFuzzer[Number]) fuzzN(numbersCount int, fuzzFunc func(t *testing.T, numbers []Number)) {
	ft.f.Helper()

	seedHandlerFunc := seedToStringAdapter(ft.f, ft.cfg.maxDigits, func(t *testing.T, strNumbers []string) {
		ft.f.Helper()

		numbers, err := parseStringSlice(ft.f, strNumbers, ft.parseNumberFunc)
		if err != nil {
			t.Errorf("failed to parse numbers: %v", err)
		}

		fuzzFunc(t, numbers)
	})

	rvFunc := createSeedFunc(ft.f, ft.cfg.signed, numbersCount, ft.cfg.uintsPerNumber, seedHandlerFunc)

	ft.f.Fuzz(rvFunc.Interface())
}

func (ft *GenericFuzzer[Number]) Fuzz1(testFunc func(t *testing.T, x1 Number)) {
	ft.f.Helper()

	ft.fuzzN(1, func(t *testing.T, numbers []Number) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0])
	})
}

func (ft *GenericFuzzer[Number]) Fuzz2(testFunc func(t *testing.T, x1, x2 Number)) {
	ft.f.Helper()

	ft.fuzzN(2, func(t *testing.T, numbers []Number) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1])
	})
}

func (ft *GenericFuzzer[Number]) Fuzz3(testFunc func(t *testing.T, x1, x2, x3 Number)) {
	ft.f.Helper()

	ft.fuzzN(3, func(t *testing.T, numbers []Number) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2])
	})
}

func (ft *GenericFuzzer[Number]) Fuzz4(testFunc func(t *testing.T, x1, x2, x3, x4 Number)) {
	ft.f.Helper()

	ft.fuzzN(4, func(t *testing.T, numbers []Number) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3])
	})
}

func (ft *GenericFuzzer[Number]) Fuzz5(testFunc func(t *testing.T, x1, x2, x3, x4, x5 Number)) {
	ft.f.Helper()

	ft.fuzzN(5, func(t *testing.T, numbers []Number) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
	})
}

func (ft *GenericFuzzer[Number]) Fuzz6(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6 Number)) {
	ft.f.Helper()

	ft.fuzzN(6, func(t *testing.T, numbers []Number) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
	})
}

func (ft *GenericFuzzer[Number]) Fuzz7(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7 Number)) {
	ft.f.Helper()

	ft.fuzzN(7, func(t *testing.T, numbers []Number) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
	})
}

func (ft *GenericFuzzer[Number]) Fuzz8(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8 Number)) {
	ft.f.Helper()

	ft.fuzzN(8, func(t *testing.T, numbers []Number) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
	})
}

func (ft *GenericFuzzer[Number]) Fuzz9(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9 Number)) {
	ft.f.Helper()

	ft.fuzzN(9, func(t *testing.T, numbers []Number) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
	})
}
func (ft *GenericFuzzer[Number]) Fuzz10(testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 Number)) {
	ft.f.Helper()

	ft.fuzzN(10, func(t *testing.T, numbers []Number) {
		t.Helper()
		ft.f.Helper()

		testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
	})
}
