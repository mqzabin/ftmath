package fuzzdecimal

import (
	"testing"
)

type GenericComparisonFuzzer[Number, Reference any] struct {
	f                        *testing.F
	parseNumberFunc          func(f *testing.F, s string) (Number, error)
	parseReferenceNumberFunc func(f *testing.F, s string) (Reference, error)
	cfg                      config
}

func NewGenericComparisonFuzzer[Number, Reference any](
	f *testing.F,
	parseNumberFunc func(f *testing.F, s string) (Number, error),
	parseReferenceNumberFunc func(f *testing.F, s string) (Reference, error),
	options ...Option,
) *GenericComparisonFuzzer[Number, Reference] {
	f.Helper()

	cfg := parseConfig(f, options)

	return &GenericComparisonFuzzer[Number, Reference]{
		f:                        f,
		parseNumberFunc:          parseNumberFunc,
		parseReferenceNumberFunc: parseReferenceNumberFunc,
		cfg:                      cfg,
	}
}

func (ft *GenericComparisonFuzzer[Number, Reference]) fuzzN(
	numbersCount int,
	refResultFunc func(f *testing.F, numbers []Reference) (Reference, error),
	fuzzFunc func(t *testing.T, refResult Reference, numbers []Number),
) {
	ft.f.Helper()

	seedHandlerFunc := seedToStringAdapter(ft.f, ft.cfg.maxDigits, func(t *testing.T, strNumbers []string) {
		ft.f.Helper()
		t.Helper()

		referenceNumbers, err := parseStringSlice(ft.f, strNumbers, ft.parseReferenceNumberFunc)
		if err != nil {
			t.Errorf("failed to parse reference numbers: %v", err)
		}

		referenceResult, err := refResultFunc(ft.f, referenceNumbers)
		if err != nil {
			t.Errorf("failed to get reference result: %v: values: %s", err, errorMessageFromValues(strNumbers))
		}

		numbers, err := parseStringSlice(ft.f, strNumbers, ft.parseNumberFunc)
		if err != nil {
			t.Errorf("failed to parse numbers: %v", err)
		}

		fuzzFunc(t, referenceResult, numbers)
	})

	rvFunc := createSeedFunc(ft.f, ft.cfg.signed, numbersCount, ft.cfg.uintsPerNumber, seedHandlerFunc)

	ft.f.Fuzz(rvFunc.Interface())
}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz1(
	refResultFunc func(f *testing.F, x1 Reference) (Reference, error),
	testFunc func(t *testing.T, refResult Reference, x1 Number),
) {
	ft.f.Helper()

	ft.fuzzN(1,
		func(f *testing.F, numbers []Reference) (Reference, error) {
			f.Helper()

			return refResultFunc(f, numbers[0])
		},
		func(t *testing.T, refResult Reference, numbers []Number) {
			ft.f.Helper()

			testFunc(t, refResult, numbers[0])
		},
	)
}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz2(
	refResultFunc func(f *testing.F, x1, x2 Reference) (Reference, error),
	testFunc func(t *testing.T, refResult Reference, x1, x2 Number),
) {
	ft.f.Helper()

	ft.fuzzN(2,
		func(f *testing.F, numbers []Reference) (Reference, error) {
			f.Helper()

			return refResultFunc(f, numbers[0], numbers[1])
		},
		func(t *testing.T, refResult Reference, numbers []Number) {
			ft.f.Helper()

			testFunc(t, refResult, numbers[0], numbers[1])
		},
	)
}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz3(
	refResultFunc func(f *testing.F, x1, x2, x3 Reference) (Reference, error),
	testFunc func(t *testing.T, refResult Reference, x1, x2, x3 Number),
) {
	ft.f.Helper()

	ft.fuzzN(3,
		func(f *testing.F, numbers []Reference) (Reference, error) {
			f.Helper()

			return refResultFunc(f, numbers[0], numbers[1], numbers[2])
		},
		func(t *testing.T, refResult Reference, numbers []Number) {
			ft.f.Helper()

			testFunc(t, refResult, numbers[0], numbers[1], numbers[2])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz4(
	refResultFunc func(f *testing.F, x1, x2, x3, x4 Reference) (Reference, error),
	testFunc func(t *testing.T, refResult Reference, x1, x2, x3, x4 Number),
) {
	ft.f.Helper()

	ft.fuzzN(4,
		func(f *testing.F, numbers []Reference) (Reference, error) {
			f.Helper()

			return refResultFunc(f, numbers[0], numbers[1], numbers[2], numbers[3])
		},
		func(t *testing.T, refResult Reference, numbers []Number) {
			ft.f.Helper()

			testFunc(t, refResult, numbers[0], numbers[1], numbers[2], numbers[3])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz5(
	refResultFunc func(f *testing.F, x1, x2, x3, x4, x5 Reference) (Reference, error),
	testFunc func(t *testing.T, refResult Reference, x1, x2, x3, x4, x5 Number),
) {
	ft.f.Helper()

	ft.fuzzN(5,
		func(f *testing.F, numbers []Reference) (Reference, error) {
			f.Helper()

			return refResultFunc(f, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
		},
		func(t *testing.T, refResult Reference, numbers []Number) {
			ft.f.Helper()

			testFunc(t, refResult, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz6(
	refResultFunc func(f *testing.F, x1, x2, x3, x4, x5, x6 Reference) (Reference, error),
	testFunc func(t *testing.T, refResult Reference, x1, x2, x3, x4, x5, x6 Number),
) {
	ft.f.Helper()

	ft.fuzzN(6,
		func(f *testing.F, numbers []Reference) (Reference, error) {
			f.Helper()

			return refResultFunc(f, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
		},
		func(t *testing.T, refResult Reference, numbers []Number) {
			ft.f.Helper()

			testFunc(t, refResult, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz7(
	refResultFunc func(f *testing.F, x1, x2, x3, x4, x5, x6, x7 Reference) (Reference, error),
	testFunc func(t *testing.T, refResult Reference, x1, x2, x3, x4, x5, x6, x7 Number),
) {
	ft.f.Helper()

	ft.fuzzN(7,
		func(f *testing.F, numbers []Reference) (Reference, error) {
			f.Helper()

			return refResultFunc(f, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
		},
		func(t *testing.T, refResult Reference, numbers []Number) {
			ft.f.Helper()

			testFunc(t, refResult, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz8(
	refResultFunc func(f *testing.F, x1, x2, x3, x4, x5, x6, x7, x8 Reference) (Reference, error),
	testFunc func(t *testing.T, refResult Reference, x1, x2, x3, x4, x5, x6, x7, x8 Number),
) {
	ft.f.Helper()

	ft.fuzzN(8,
		func(f *testing.F, numbers []Reference) (Reference, error) {
			f.Helper()

			return refResultFunc(f, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
		},
		func(t *testing.T, refResult Reference, numbers []Number) {
			ft.f.Helper()

			testFunc(t, refResult, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz9(
	refResultFunc func(f *testing.F, x1, x2, x3, x4, x5, x6, x7, x8, x9 Reference) (Reference, error),
	testFunc func(t *testing.T, refResult Reference, x1, x2, x3, x4, x5, x6, x7, x8, x9 Number),
) {
	ft.f.Helper()

	ft.fuzzN(9,
		func(f *testing.F, numbers []Reference) (Reference, error) {
			f.Helper()

			return refResultFunc(f, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
		},
		func(t *testing.T, refResult Reference, numbers []Number) {
			ft.f.Helper()

			testFunc(t, refResult, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
		},
	)
}
func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz10(
	refResultFunc func(f *testing.F, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 Reference) (Reference, error),
	testFunc func(t *testing.T, refResult Reference, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 Number),
) {
	ft.fuzzN(9,
		func(f *testing.F, numbers []Reference) (Reference, error) {
			f.Helper()

			return refResultFunc(f, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
		},
		func(t *testing.T, refResult Reference, numbers []Number) {
			ft.f.Helper()

			testFunc(t, refResult, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
		},
	)
}
