package fuzzdecimal

import (
	"testing"
)

type GenericComparisonFuzzer[Number, Reference any] struct {
	f                        *testing.F
	parseNumberFunc          func(t *testing.T, s string) (Number, error)
	parseReferenceNumberFunc func(t *testing.T, s string) (Reference, error)
	cfg                      config
}

func NewGenericComparisonFuzzer[Number, Reference any](
	f *testing.F,
	parseNumberFunc func(t *testing.T, s string) (Number, error),
	parseReferenceNumberFunc func(t *testing.T, s string) (Reference, error),
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
	refResultFunc func(t *testing.T, numbers []Reference) (string, error),
	fuzzFunc func(t *testing.T, numbers []Number) string,
) {
	seedHandlerFunc := seedToStringAdapter(ft.cfg.maxDigits, func(t *testing.T, strNumbers []string) {
		t.Helper()

		referenceNumbers, err := parseStringSlice(t, strNumbers, ft.parseReferenceNumberFunc)
		if err != nil {
			t.Errorf("failed to parse reference numbers: %v", err)
		}

		referenceResultStr, err := refResultFunc(t, referenceNumbers)
		if err != nil {
			t.Errorf("failed to get reference result: %v: values:\n%s", err, errorMessageFromValues("\t", strNumbers))
		}

		numbers, err := parseStringSlice(t, strNumbers, ft.parseNumberFunc)
		if err != nil {
			t.Errorf("failed to parse numbers: %v", err)
		}

		resultStr := fuzzFunc(t, numbers)

		if resultStr != referenceResultStr {
			t.Errorf("unexpected result:\n\tgot: \n\t\t%s\n\twant: \n\t\t%s\n\tvalues:\n%s", resultStr, referenceResultStr, errorMessageFromValues("\t\t", strNumbers))
		}
	})

	rvFunc := createSeedFunc(ft.f, ft.cfg.signed, numbersCount, ft.cfg.uintsPerNumber, seedHandlerFunc)

	ft.f.Fuzz(rvFunc.Interface())
}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz1(
	refResultFunc func(t *testing.T, x1 Reference) (string, error),
	testFunc func(t *testing.T, x1 Number) string,
) {
	ft.fuzzN(1,
		func(t *testing.T, numbers []Reference) (string, error) {
			t.Helper()

			return refResultFunc(t, numbers[0])
		},
		func(t *testing.T, numbers []Number) string {
			return testFunc(t, numbers[0])
		},
	)
}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz2(
	refResultFunc func(t *testing.T, x1, x2 Reference) (string, error),
	testFunc func(t *testing.T, x1, x2 Number) string,
) {
	ft.fuzzN(2,
		func(t *testing.T, numbers []Reference) (string, error) {
			t.Helper()

			return refResultFunc(t, numbers[0], numbers[1])
		},
		func(t *testing.T, numbers []Number) string {
			return testFunc(t, numbers[0], numbers[1])
		},
	)
}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz3(
	refResultFunc func(t *testing.T, x1, x2, x3 Reference) (string, error),
	testFunc func(t *testing.T, x1, x2, x3 Number) string,
) {
	ft.fuzzN(3,
		func(t *testing.T, numbers []Reference) (string, error) {
			t.Helper()

			return refResultFunc(t, numbers[0], numbers[1], numbers[2])
		},
		func(t *testing.T, numbers []Number) string {
			return testFunc(t, numbers[0], numbers[1], numbers[2])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz4(
	refResultFunc func(t *testing.T, x1, x2, x3, x4 Reference) (string, error),
	testFunc func(t *testing.T, x1, x2, x3, x4 Number) string,
) {
	ft.fuzzN(4,
		func(t *testing.T, numbers []Reference) (string, error) {
			t.Helper()

			return refResultFunc(t, numbers[0], numbers[1], numbers[2], numbers[3])
		},
		func(t *testing.T, numbers []Number) string {
			return testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz5(
	refResultFunc func(t *testing.T, x1, x2, x3, x4, x5 Reference) (string, error),
	testFunc func(t *testing.T, x1, x2, x3, x4, x5 Number) string,
) {
	ft.fuzzN(5,
		func(t *testing.T, numbers []Reference) (string, error) {
			t.Helper()

			return refResultFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
		},
		func(t *testing.T, numbers []Number) string {
			return testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz6(
	refResultFunc func(t *testing.T, x1, x2, x3, x4, x5, x6 Reference) (string, error),
	testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6 Number) string,
) {
	ft.fuzzN(6,
		func(t *testing.T, numbers []Reference) (string, error) {
			t.Helper()

			return refResultFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
		},
		func(t *testing.T, numbers []Number) string {
			return testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz7(
	refResultFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7 Reference) (string, error),
	testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7 Number) string,
) {
	ft.fuzzN(7,
		func(t *testing.T, numbers []Reference) (string, error) {
			t.Helper()

			return refResultFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
		},
		func(t *testing.T, numbers []Number) string {
			return testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz8(
	refResultFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8 Reference) (string, error),
	testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8 Number) string,
) {
	ft.fuzzN(8,
		func(t *testing.T, numbers []Reference) (string, error) {
			t.Helper()

			return refResultFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
		},
		func(t *testing.T, numbers []Number) string {
			return testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
		},
	)

}

func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz9(
	refResultFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9 Reference) (string, error),
	testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9 Number) string,
) {
	ft.fuzzN(9,
		func(t *testing.T, numbers []Reference) (string, error) {
			t.Helper()

			return refResultFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
		},
		func(t *testing.T, numbers []Number) string {
			return testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
		},
	)
}
func (ft *GenericComparisonFuzzer[Number, Reference]) Fuzz10(
	refResultFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 Reference) (string, error),
	testFunc func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 Number) string,
) {
	ft.fuzzN(9,
		func(t *testing.T, numbers []Reference) (string, error) {
			t.Helper()

			return refResultFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
		},
		func(t *testing.T, numbers []Number) string {
			return testFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
		},
	)
}
