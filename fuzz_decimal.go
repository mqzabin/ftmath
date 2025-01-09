package fuzzdecimal

import "github.com/mqzabin/fuzzdecimal/fdlib"

const asDecimalFuncName = "AsDecimal"

func AsDecimalSlice[Decimal any](t *T, name string, parseDecimalFunc func(t *T, s string) (Decimal, error), fuzzFunc func(t *T, numbers []Decimal)) {
	t.Helper()

	parseDecimal := convertParseDecimalFunc(t, parseDecimalFunc)

	t.Run(name, func(t *T) {
		decimals, err := fdlib.ParseStringSliceToDecimalSlice(t.T, t.seeds, parseDecimal)
		if err != nil {
			t.Fatalf("failed to parse decimals: %v", err)
		}

		fuzzFunc(t, decimals)
	})
}

func AsDecimal1[Decimal any](t *T, name string, parseDecimalFunc func(t *T, s string) (Decimal, error), fuzzFunc func(t *T, x1 Decimal)) {
	t.Helper()

	AsDecimalSlice(t, name, parseDecimalFunc, func(t *T, numbers []Decimal) {
		t.Helper()

		t.assertStaticSeedsCount(1, asDecimalFuncName)

		fuzzFunc(t, numbers[0])
	})
}

func AsDecimal2[Decimal any](t *T, name string, parseDecimalFunc func(t *T, s string) (Decimal, error), fuzzFunc func(t *T, x1, x2 Decimal)) {
	t.Helper()

	AsDecimalSlice(t, name, parseDecimalFunc, func(t *T, numbers []Decimal) {
		t.Helper()

		t.assertStaticSeedsCount(2, asDecimalFuncName)

		fuzzFunc(t, numbers[0], numbers[1])
	})
}

func AsDecimal3[Decimal any](t *T, name string, parseDecimalFunc func(t *T, s string) (Decimal, error), fuzzFunc func(t *T, x1, x2, x3 Decimal)) {
	t.Helper()

	AsDecimalSlice(t, name, parseDecimalFunc, func(t *T, numbers []Decimal) {
		t.Helper()

		t.assertStaticSeedsCount(3, asDecimalFuncName)

		fuzzFunc(t, numbers[0], numbers[1], numbers[2])
	})
}

func AsDecimal4[Decimal any](t *T, name string, parseDecimalFunc func(t *T, s string) (Decimal, error), fuzzFunc func(t *T, x1, x2, x3, x4 Decimal)) {
	t.Helper()

	AsDecimalSlice(t, name, parseDecimalFunc, func(t *T, numbers []Decimal) {
		t.Helper()

		t.assertStaticSeedsCount(4, asDecimalFuncName)

		fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3])
	})
}

func AsDecimal5[Decimal any](t *T, name string, parseDecimalFunc func(t *T, s string) (Decimal, error), fuzzFunc func(t *T, x1, x2, x3, x4, x5 Decimal)) {
	t.Helper()

	AsDecimalSlice(t, name, parseDecimalFunc, func(t *T, numbers []Decimal) {
		t.Helper()

		t.assertStaticSeedsCount(5, asDecimalFuncName)

		fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
	})
}

func AsDecimal6[Decimal any](t *T, name string, parseDecimalFunc func(t *T, s string) (Decimal, error), fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6 Decimal)) {
	t.Helper()

	AsDecimalSlice(t, name, parseDecimalFunc, func(t *T, numbers []Decimal) {
		t.Helper()

		t.assertStaticSeedsCount(6, asDecimalFuncName)

		fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
	})
}

func AsDecimal7[Decimal any](t *T, name string, parseDecimalFunc func(t *T, s string) (Decimal, error), fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7 Decimal)) {
	t.Helper()

	AsDecimalSlice(t, name, parseDecimalFunc, func(t *T, numbers []Decimal) {
		t.Helper()

		t.assertStaticSeedsCount(7, asDecimalFuncName)

		fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
	})
}

func AsDecimal8[Decimal any](t *T, name string, parseDecimalFunc func(t *T, s string) (Decimal, error), fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7, x8 Decimal)) {
	t.Helper()

	AsDecimalSlice(t, name, parseDecimalFunc, func(t *T, numbers []Decimal) {
		t.Helper()

		t.assertStaticSeedsCount(8, asDecimalFuncName)

		fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
	})
}

func AsDecimal9[Decimal any](t *T, name string, parseDecimalFunc func(t *T, s string) (Decimal, error), fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7, x8, x9 Decimal)) {
	t.Helper()

	AsDecimalSlice(t, name, parseDecimalFunc, func(t *T, numbers []Decimal) {
		t.Helper()

		t.assertStaticSeedsCount(9, asDecimalFuncName)

		fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
	})
}

func AsDecimal10[Decimal any](t *T, name string, parseDecimalFunc func(t *T, s string) (Decimal, error), fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 Decimal)) {
	t.Helper()

	AsDecimalSlice(t, name, parseDecimalFunc, func(t *T, numbers []Decimal) {
		t.Helper()

		t.assertStaticSeedsCount(10, asDecimalFuncName)

		fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
	})
}
