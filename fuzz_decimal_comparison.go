package fuzzdecimal

const asDecimalComparisonFuncName = "AsDecimalComparison"

func AsDecimalComparisonSlice[Decimal, ReferenceDecimal any](
	t *T, name string,
	parseDecimalFunc func(t *T, s string) (Decimal, error),
	parseReferenceDecimalFunc func(t *T, s string) (ReferenceDecimal, error),
	computeReferenceResult func(t *T, numbers []ReferenceDecimal) (string, error),
	fuzzFunc func(t *T, numbers []Decimal) string,
) {
	t.Helper()

	t.Run(name, func(t *T) {
		refDecimals, err := parseStringSlice(t, t.seeds, parseReferenceDecimalFunc)
		if err != nil {
			t.Errorf("failed to parse reference decimals: %v", err)
		}

		refResult, err := computeReferenceResult(t, refDecimals)
		if err != nil {
			t.Errorf("failed to get reference result: %v: values:\n%s", err, t.seedsErrorMessage("\t"))
		}

		decimals, err := parseStringSlice(t, t.seeds, parseDecimalFunc)
		if err != nil {
			t.Errorf("failed to parse decimals: %v", err)
		}

		result := fuzzFunc(t, decimals)

		if result != refResult {
			t.Errorf("unexpected result:\n\tgot: \n\t\t%s\n\twant: \n\t\t%s\n\tvalues:\n%s", result, refResult, t.seedsErrorMessage("\t\t"))
		}
	})
}

func AsDecimalComparison1[Decimal, ReferenceDecimal any](
	t *T, name string,
	parseDecimalFunc func(t *T, s string) (Decimal, error),
	parseReferenceDecimalFunc func(t *T, s string) (ReferenceDecimal, error),
	computeReferenceResult func(t *T, x1 ReferenceDecimal) (string, error),
	fuzzFunc func(t *T, x1 Decimal) string,
) {
	t.Helper()

	AsDecimalComparisonSlice(t, name, parseDecimalFunc, parseReferenceDecimalFunc,
		func(t *T, numbers []ReferenceDecimal) (string, error) {
			t.Helper()

			t.assertStaticSeedsCount(1, asDecimalComparisonFuncName)

			return computeReferenceResult(t, numbers[0])
		},
		func(t *T, numbers []Decimal) string {
			t.Helper()

			return fuzzFunc(t, numbers[0])
		})
}

func AsDecimalComparison2[Decimal, ReferenceDecimal any](
	t *T,
	name string,
	parseDecimalFunc func(t *T, s string) (Decimal, error),
	parseReferenceDecimalFunc func(t *T, s string) (ReferenceDecimal, error),
	computeReferenceResult func(t *T, x1, x2 ReferenceDecimal) (string, error),
	fuzzFunc func(t *T, x1, x2 Decimal) string,
) {
	t.Helper()

	AsDecimalComparisonSlice(t, name, parseDecimalFunc, parseReferenceDecimalFunc,
		func(t *T, numbers []ReferenceDecimal) (string, error) {
			t.Helper()

			t.assertStaticSeedsCount(2, asDecimalComparisonFuncName)

			return computeReferenceResult(t, numbers[0], numbers[1])
		},
		func(t *T, numbers []Decimal) string {
			t.Helper()

			return fuzzFunc(t, numbers[0], numbers[1])
		})
}

func AsDecimalComparison3[Decimal, ReferenceDecimal any](
	t *T,
	name string,
	parseDecimalFunc func(t *T, s string) (Decimal, error),
	parseReferenceDecimalFunc func(t *T, s string) (ReferenceDecimal, error),
	computeReferenceResult func(t *T, x1, x2, x3 ReferenceDecimal) (string, error),
	fuzzFunc func(t *T, x1, x2, x3 Decimal) string,
) {
	t.Helper()

	AsDecimalComparisonSlice(t, name, parseDecimalFunc, parseReferenceDecimalFunc,
		func(t *T, numbers []ReferenceDecimal) (string, error) {
			t.Helper()

			t.assertStaticSeedsCount(3, asDecimalComparisonFuncName)

			return computeReferenceResult(t, numbers[0], numbers[1], numbers[2])
		},
		func(t *T, numbers []Decimal) string {
			t.Helper()

			return fuzzFunc(t, numbers[0], numbers[1], numbers[2])
		})
}

func AsDecimalComparison4[Decimal, ReferenceDecimal any](
	t *T,
	name string,
	parseDecimalFunc func(t *T, s string) (Decimal, error),
	parseReferenceDecimalFunc func(t *T, s string) (ReferenceDecimal, error),
	computeReferenceResult func(t *T, x1, x2, x3, x4 ReferenceDecimal) (string, error),
	fuzzFunc func(t *T, x1, x2, x3, x4 Decimal) string,
) {
	t.Helper()

	AsDecimalComparisonSlice(t, name, parseDecimalFunc, parseReferenceDecimalFunc,
		func(t *T, numbers []ReferenceDecimal) (string, error) {
			t.Helper()

			t.assertStaticSeedsCount(4, asDecimalComparisonFuncName)

			return computeReferenceResult(t, numbers[0], numbers[1], numbers[2], numbers[3])
		},
		func(t *T, numbers []Decimal) string {
			t.Helper()

			return fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3])
		})
}

func AsDecimalComparison5[Decimal, ReferenceDecimal any](
	t *T,
	name string,
	parseDecimalFunc func(t *T, s string) (Decimal, error),
	parseReferenceDecimalFunc func(t *T, s string) (ReferenceDecimal, error),
	computeReferenceResult func(t *T, x1, x2, x3, x4, x5 ReferenceDecimal) (string, error),
	fuzzFunc func(t *T, x1, x2, x3, x4, x5 Decimal) string,
) {
	t.Helper()

	AsDecimalComparisonSlice(t, name, parseDecimalFunc, parseReferenceDecimalFunc,
		func(t *T, numbers []ReferenceDecimal) (string, error) {
			t.Helper()

			t.assertStaticSeedsCount(5, asDecimalComparisonFuncName)

			return computeReferenceResult(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
		},
		func(t *T, numbers []Decimal) string {
			t.Helper()

			return fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
		})
}

func AsDecimalComparison6[Decimal, ReferenceDecimal any](
	t *T,
	name string,
	parseDecimalFunc func(t *T, s string) (Decimal, error),
	parseReferenceDecimalFunc func(t *T, s string) (ReferenceDecimal, error),
	computeReferenceResult func(t *T, x1, x2, x3, x4, x5, x6 ReferenceDecimal) (string, error),
	fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6 Decimal) string,
) {
	t.Helper()

	AsDecimalComparisonSlice(t, name, parseDecimalFunc, parseReferenceDecimalFunc,
		func(t *T, numbers []ReferenceDecimal) (string, error) {
			t.Helper()

			t.assertStaticSeedsCount(6, asDecimalComparisonFuncName)

			return computeReferenceResult(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
		},
		func(t *T, numbers []Decimal) string {
			t.Helper()

			return fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
		})
}

func AsDecimalComparison7[Decimal, ReferenceDecimal any](
	t *T,
	name string,
	parseDecimalFunc func(t *T, s string) (Decimal, error),
	parseReferenceDecimalFunc func(t *T, s string) (ReferenceDecimal, error),
	computeReferenceResult func(t *T, x1, x2, x3, x4, x5, x6, x7 ReferenceDecimal) (string, error),
	fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7 Decimal) string,
) {
	t.Helper()

	AsDecimalComparisonSlice(t, name, parseDecimalFunc, parseReferenceDecimalFunc,
		func(t *T, numbers []ReferenceDecimal) (string, error) {
			t.Helper()

			t.assertStaticSeedsCount(7, asDecimalComparisonFuncName)

			return computeReferenceResult(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
		},
		func(t *T, numbers []Decimal) string {
			t.Helper()

			return fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
		})
}

func AsDecimalComparison8[Decimal, ReferenceDecimal any](
	t *T,
	name string,
	parseDecimalFunc func(t *T, s string) (Decimal, error),
	parseReferenceDecimalFunc func(t *T, s string) (ReferenceDecimal, error),
	computeReferenceResult func(t *T, x1, x2, x3, x4, x5, x6, x7, x8 ReferenceDecimal) (string, error),
	fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7, x8 Decimal) string,
) {
	t.Helper()

	AsDecimalComparisonSlice(t, name, parseDecimalFunc, parseReferenceDecimalFunc,
		func(t *T, numbers []ReferenceDecimal) (string, error) {
			t.Helper()

			t.assertStaticSeedsCount(8, asDecimalComparisonFuncName)

			return computeReferenceResult(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
		},
		func(t *T, numbers []Decimal) string {
			t.Helper()

			return fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
		})
}

func AsDecimalComparison9[Decimal, ReferenceDecimal any](
	t *T,
	name string,
	parseDecimalFunc func(t *T, s string) (Decimal, error),
	parseReferenceDecimalFunc func(t *T, s string) (ReferenceDecimal, error),
	computeReferenceResult func(t *T, x1, x2, x3, x4, x5, x6, x7, x8, x9 ReferenceDecimal) (string, error),
	fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7, x8, x9 Decimal) string,
) {
	t.Helper()

	AsDecimalComparisonSlice(t, name, parseDecimalFunc, parseReferenceDecimalFunc,
		func(t *T, numbers []ReferenceDecimal) (string, error) {
			t.Helper()

			t.assertStaticSeedsCount(9, asDecimalComparisonFuncName)

			return computeReferenceResult(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
		},
		func(t *T, numbers []Decimal) string {
			t.Helper()

			return fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
		})
}

func AsDecimalComparison10[Decimal, ReferenceDecimal any](
	t *T,
	name string,
	parseDecimalFunc func(t *T, s string) (Decimal, error),
	parseReferenceDecimalFunc func(t *T, s string) (ReferenceDecimal, error),
	computeReferenceResult func(t *T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 ReferenceDecimal) (string, error),
	fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 Decimal) string,
) {
	t.Helper()

	AsDecimalComparisonSlice(t, name, parseDecimalFunc, parseReferenceDecimalFunc,
		func(t *T, numbers []ReferenceDecimal) (string, error) {
			t.Helper()

			t.assertStaticSeedsCount(10, asDecimalComparisonFuncName)

			return computeReferenceResult(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
		},
		func(t *T, numbers []Decimal) string {
			t.Helper()

			return fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
		})
}
